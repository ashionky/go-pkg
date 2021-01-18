package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ivpusic/grpool"
)

type TestEventHandler struct {
	P *grpool.Pool
}

func (handler TestEventHandler) HandleKafkaMsg(message *ReportEvent) (err error) {
	if  checkTestRepeatEvent(message.Id) {
		err = errors.New("event msg repeat id: " + message.Id)
		return err
	}
	switch message.Type {
	case "test":
	   //todo 处理数据message。。。。
	   fmt.Println("body==",message.Body)

		CopyTestEvent(message)  //记录已处理的消息
	default:
		fmt.Println("can not process event:" + message.Type)
	}
	return
}

func (handler TestEventHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (handler TestEventHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (handler TestEventHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		eventmsg := ReportEvent{}
		err := json.Unmarshal(msg.Value, &eventmsg)
		if err != nil {
			fmt.Printf("parse kafka message error: %v \n", err)
		}
		handler.HandleKafkaMsg(&eventmsg)

		fmt.Printf("finish===============%p \n", &handler)
		sess.MarkMessage(msg, "")
	}
	return nil
}

// 检查事件是否重复，若重复则忽略，处理事件之前调用
func checkTestRepeatEvent(id string) bool {
	return false
}

/*
记录已处理事件，用于消费事件的时候过滤重复事件
*/
func CopyTestEvent(eventData *ReportEvent) {


}
