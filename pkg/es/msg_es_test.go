/**
 * @Author pibing
 * @create 2021/1/13 12:18 PM
 */

package es

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"
	"time"
)

func TestAddMessage(t *testing.T) {
	var configFile  ="../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	Init()
	mp :=map[string]interface{}{
		"_id":"112",
		"name":"张sss",
		"age":20,
		"address":"成都的",
	}
	AddMessage(mp,"a")
	fmt.Println("=====")
    time.Sleep(5*time.Second)
	data, err := SearchMessageAll("a", "张")
	fmt.Println("=",data)
	fmt.Println("==",err)

}

func TestSearchMessageAll(t *testing.T) {



}
