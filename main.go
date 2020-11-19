/**
 * @Author pibing
 * @create 2020/11/14 12:48 PM
 */

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pkg/model"
	"go-pkg/pkg/cfg"
	"go-pkg/pkg/db"
	"go-pkg/pkg/log"
	"go-pkg/pkg/mongodb"
	"go-pkg/pkg/redis"
	"go-pkg/router"
	"os"
)


var config=cfg.GetConfig()

func main()  {
	mode := flag.String("m", "dev", "指定执行模式,只支持[dev|test|prod],默认是dev")
	flag.Parse()
	dev := true
	if *mode != "dev" {
		dev = false
	}
	if dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
   //初始化config
	configFile := fmt.Sprintf("./conf/%s.yml", *mode)
	err := cfg.Initcfg(configFile)
	if err != nil {
		fmt.Printf("读取配置文件失败[%s]: %s\n", configFile, err.Error())
		os.Exit(1)
	}

	//初始化日志工具
	log.InitLog()

    //初始化数据库连接
	err = InitDB()
	if err != nil {
		fmt.Printf("数据库初始化失败: %s\n", err.Error())
		log.Info("数据库初始化失败", err.Error())
		os.Exit(2)
	}

	//mongo初始化
	err =mongodb.MongoInit()
	if err != nil {
		fmt.Printf("mongodb初始化失败: %s\n", err.Error())
		log.Info("mongodb初始化失败", err.Error())
		os.Exit(3)
	}

    //初始化redis
	err = InitRedis()
	if err != nil {
		fmt.Printf("redis初始化失败: %s\n", err.Error())
		log.Info("redis初始化失败", err.Error())
		os.Exit(4)
	}

    //初始化路由组
	err = router.InitRouter()
	if err != nil {
		fmt.Printf("服务器启动失败: %s\n", err.Error())
		log.Info("服务器启动失败", err.Error())
		os.Exit(5)
	}

	fmt.Println("程序已启动")

	// 阻塞
	select {}

}

// 连接数据库
func InitDB() (err error) {
	driver := "mysql"

	user := config.Mysql.User
	password := config.Mysql.Password
	host := config.Mysql.Host
	dbname := config.Mysql.Dbname
	charset := config.Mysql.Charset
	sqlconnStr := fmt.Sprintf("%v:%v@(%v)/%v?charset=%v&parseTime=True&loc=Local",
		user, password, host, dbname, charset)
	_, err = db.InitDefaultDB(sqlconnStr, driver, nil)
	//自动创建表结构
	model.InitTables()
	return err
}

// 连接redis数据库
func InitRedis() error {
	host := config.Redis.Host
	auth := config.Redis.Password
	rdb := config.Redis.Db
	maxActive := config.Redis.Max_active
	maxIdle := config.Redis.Max_idle
	idleTimeout := config.Redis.Idle_timeout

	err := redis.Init(host, auth, rdb, maxActive, maxIdle, idleTimeout)
	return err
}
