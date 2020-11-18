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

func TestInit(t *testing.T) {
	var configFile  ="../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	var config=cfg.GetConfig()

	host := config.Redis.Host
	auth := config.Redis.Password
	rdb := config.Redis.Db
	maxActive := config.Redis.Max_active
	maxIdle := config.Redis.Max_idle
	idleTimeout := config.Redis.Idle_timeout
	_ = Init(host, auth, rdb, maxActive, maxIdle, idleTimeout)

	defaultRedis.Set("test","88888")
	b,_:=defaultRedis.Get("test")
	fmt.Print(string(b))
}
