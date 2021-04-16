/**
 * @Author pibing
 * @create 2020/11/15 12:54 PM
 */

package kafka

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"go-pkg/pkg/util"
	"testing"
)

func TestConsumer(t *testing.T) {

}

func TestSendEvent(t *testing.T) {
	//加载config
	var configFile = "../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	Init()

	//发送消息
	eventDataCus := ReportEvent{
		Id:      util.UUID(),
		Topic:   "test",
		Type:    "test",
		Account: "123",
		Time:    util.GetNowDateTimeFormat(),
		Body:    map[string]interface{}{"name": "张三aaa"},
	}
	err := SendEvent(&eventDataCus)
	fmt.Println("err:", err)
	fmt.Println("sent success")
}
