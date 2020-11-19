/**
 * @Author pibing
 * @create 2020/11/16 9:55 AM
 */

package mongodb

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"
)

func TestMongoInit(t *testing.T) {
	var configFile  ="../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	MongoInit()
    mp:=map[string]interface{}{
    	"name":"lisi",
    	"age":10,
	}
	DbInsert("test","user",mp)
    map1:=make(map[string]interface{})
    DbFindOne("test","user",B{},&map1)
    fmt.Print(map1)

}
