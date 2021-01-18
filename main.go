/**
 * @Author pibing
 * @create 2020/11/14 12:48 PM
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pkg/model"
	"go-pkg/pkg/cfg"
	"go-pkg/pkg/db"
	"go-pkg/pkg/kafka"
	"go-pkg/pkg/log"
	"go-pkg/pkg/mongodb"
	"go-pkg/pkg/redis"
	"go-pkg/router"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
	err =InitMongodb()
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
	r := router.InitRouter()

	// kafka的上下文
	kafka.Init()
	ctx1, cancel1 := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	kafka.InitHandler(ctx1, wg)

	// 服务优雅退出
	srv := http.Server{
		Addr:    ":" + config.Server.Http_port,
		Handler: r,
	}
	log.Info("ListenAndServe port:", config.Server.Http_port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen:", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:", err.Error())
	}
	cancel1()
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

	err := redis.Init(host, auth, rdb, maxActive)
	return err
}

//mongodb连接
func InitMongodb() error {
	var url = "mongodb://" + config.Mongodb.User + ":" + config.Mongodb.Password + "@" + config.Mongodb.Host + ":" + config.Mongodb.Port + "/admin"
	err :=mongodb.MongoInit(url,config.Mongodb.Dbname,config.Mongodb.Poolsize)
	return err
}
