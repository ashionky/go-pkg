/**
 * @Author pibing
 * @create 2020/11/16 9:55 AM
 */

package mongodb

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"
	"time"
)

func TestMongoInit(t *testing.T) {
	var configFile = "../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	config := cfg.GetConfig()
	var url = "mongodb://" + config.Mongodb.User + ":" + config.Mongodb.Password + "@" + config.Mongodb.Host + ":" + config.Mongodb.Port + "/admin"
	MongoInit(url, config.Mongodb.Dbname, config.Mongodb.Poolsize)
	mp := map[string]interface{}{
		"name": "lisi",
		"age":  10,
		"createdAt":time.Now(),
	}

	CreateIndex("test", "user",  "createdAt",  60  )


	DbInsert("test", "user", mp)
	fmt.Println("=1=",time.Now())
	//map1 := make(map[string]interface{})
	//DbFindOne("test", "user", B{}, &map1)
	//fmt.Print(map1)

	time.Sleep(50*time.Second)
	mp2 := B{
		"name": "qq",
		"field1":"beijing",
		"createdAt":time.Now(),
	}

	err := DbUpdateOne("test", "user", B{"age": 10}, mp2).Error()
	fmt.Println("=2=",time.Now())
	fmt.Println("err=",err)

}


