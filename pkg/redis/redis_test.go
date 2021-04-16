/**
 * @Author pibing
 * @create 2020/11/16 9:50 AM
 */

package redis

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"
)

type user struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func TestInit(t *testing.T) {
	var configFile = "../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	var config = cfg.GetConfig()

	host := config.Redis.Host
	auth := config.Redis.Password
	rdb := config.Redis.Db
	maxActive := config.Redis.Max_active

	_ = Init(host, auth, rdb, maxActive)

	defaultRedis.Set("user", "info")
	b, err := defaultRedis.Get("user")
	//var mp2  user
	//json.Unmarshal(b,&mp2)
	fmt.Println(err)
	fmt.Print(string(b))
}
