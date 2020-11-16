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
    	"name":"test",
    	"age":10,
	}
	Insert("test_mo",mp)
    map1:=make(map[string]interface{})
    FindOne("test_mo",B{},&map1)
    fmt.Print(map1)

}
