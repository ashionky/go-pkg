/**
 * @Author pibing
 * @create 2021/1/3 12:43 PM
 */

package go_redis

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"
)
type user struct {
	Name string `json:"name"`
	Address string  `json:"address"`
}
func TestMG( t *testing.T)  {
	var configFile  ="../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	var config=cfg.GetConfig()

	host := config.Redis.Host
	auth := config.Redis.Password
	rdb := config.Redis.Db
	maxActive := config.Redis.Max_active

	_ = RedisInit(host, auth, rdb, maxActive)


	HSet("user","age",10)
	HSet("user","name","ooooo")
	var mp2  string
	_=HGet("user","name",&mp2)
	fmt.Print(mp2)
	Hdel("user","name")

	//json.Unmarshal([]byte(b),&mp2)
	err :=HGet("user","name",&mp2)
	fmt.Println(err)
	fmt.Print(mp2)




}
