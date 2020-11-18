package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"go-pkg/pkg/log"
	"go-pkg/pkg/cfg"
	"math"
	"strings"
	"time"
)

//var job *grpool.Pool

//最大批量推送消息数量
const maxBatchNum int = 500

var kfkConfig *sarama.Config
var kfkOrderConfig *sarama.Config
var Addresses []string
var groupID string
var config  = cfg.GetConfig()

func getKafkaConfig() *sarama.Config {
	if kfkConfig == nil {
		kfkConfig = sarama.NewConfig()
		kfkConfig.Producer.Return.Successes = true
		kfkConfig.Version = sarama.V1_0_0_0
		kfkConfig.Producer.MaxMessageBytes = 10000000
	}
	return kfkConfig
}

func getKafkOrderConfig() *sarama.Config {
	if kfkOrderConfig == nil {
		kfkOrderConfig = sarama.NewConfig()
		kfkOrderConfig.Net.MaxOpenRequests = 1 //正在发送但是发送状态未知的最大消息数量，设为1，解决重试发送导致的消息乱序问题
		kfkOrderConfig.Producer.Return.Successes = true
		kfkOrderConfig.Version = sarama.V1_0_0_0
		kfkConfig.Producer.MaxMessageBytes = 10000000
	}
	return kfkOrderConfig
}

func Init() {
	//groupID = ""
	ips := config.Kafka.Url
	for _, ip := range strings.Split(ips, ",") {
		Addresses = append(Addresses, ip)
	}
	log.Info("kafka brokers=====", ips)
}

/**
发送无顺序的单条事件
@param
eventData: 事件对象
*/
func SendEvent(eventData *ReportEvent) error {
	if eventData.RetryTime == "" && eventData.Num == 0 {
		eventData.Num = 1
	} else {
		eventData.Num = eventData.Num + 1
	}
	eventData.RetryTime = GetRetryTime(eventData.Num)

	if eventData.Topic == "" {
		return errors.New("topic is empty")
	}
	if eventData.Id == "" {
		return errors.New("eventData is empty")
	}
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		log.Info("json marshal err: %+v %v", eventData, err)
		return err
	}

	p, err := sarama.NewSyncProducer(Addresses, getKafkaConfig())
	if err != nil {
		log.Info("sarama.NewSyncProducer err, message=%s", err)
		go recordFailEvent(eventData)
		return err
	}
	defer p.Close()

	msg := &sarama.ProducerMessage{
		Topic: eventData.Topic,
		Value: sarama.ByteEncoder(jsonData),
	}
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Info("publish event err: %v", err)
		go recordFailEvent(eventData)
		return err
	} else {
		log.Info(string(jsonData)+"send message success，partition=%d, offset=%d \n", part, offset)
	}
	return nil
}

/**
发送有顺序的单条事件,需要用固定的分区
@param
eventData: 事件对象
*/
func SendOrderEvent(eventData *ReportEvent) error {
	if eventData.RetryTime == "" && eventData.Num == 0 {
		eventData.Num = 1
	} else {
		eventData.Num = eventData.Num + 1
	}
	eventData.RetryTime = GetRetryTime(eventData.Num)

	if eventData.Topic == "" {
		return errors.New("topic is empty")
	}
	if eventData.Id == "" {
		return errors.New("eventData is empty")
	}
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		log.Info("json marshal err: %+v %v", eventData, err)
		return err
	}

	p, err := sarama.NewSyncProducer(Addresses, getKafkOrderConfig())
	if err != nil {
		log.Info("sarama.NewSyncProducer err, message=%s", err)
		go recordFailEvent(eventData)
		return err
	}
	defer p.Close()

	msg := &sarama.ProducerMessage{
		Topic: eventData.Topic,
		Value: sarama.ByteEncoder(jsonData),
		Key:   sarama.ByteEncoder(eventData.Account), //根据账号Id路由分区
	}
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Info("publish event err: %v", err)
		go recordFailEvent(eventData)
		return err
	} else {
		log.Info("msgId:"+eventData.Id+",send message success，partition=%d, offset=%d \n", part, offset)
	}
	return nil
}

/**
发送无顺序的批量事件
@param
topic： 主题
eventList: 事件列表
*/
func SendBatchEvent(topic string, eventList []ReportEvent) (err error) {
	if topic == "" {
		return errors.New("topic is empty")
	}
	if len(eventList) <= 0 {
		return errors.New("eventList is empty")
	}
	listSize := len(eventList)
	f := float64(listSize) / float64(maxBatchNum)
	num := int(math.Ceil(f))
	start := 0
	end := maxBatchNum
	for i := 0; i < num; i++ {
		if i == num-1 {
			end = listSize
		}
		subList := eventList[start:end]
		err = publishMultiEvent(topic, subList, false)
		if err != nil {
			recordBatchFailEvent(subList)
		}
		start += maxBatchNum
		end += maxBatchNum
	}
	return
}

/**
发送有顺序的批量事件
@param
topic： 主题
eventList: 事件列表
*/
func SendBatchOrderEvent(topic string, eventList []ReportEvent) (err error) {
	if topic == "" {
		return errors.New("topic is empty")
	}
	if len(eventList) <= 0 {
		return errors.New("eventList is empty")
	}
	listSize := len(eventList)
	f := float64(listSize) / float64(maxBatchNum)
	num := int(math.Ceil(f))
	start := 0
	end := maxBatchNum
	for i := 0; i < num; i++ {
		if i == num-1 {
			end = listSize
		}
		subList := eventList[start:end]
		err = publishMultiEvent(topic, subList, true)
		if err != nil {
			recordBatchFailEvent(subList)
		}
		start += maxBatchNum
		end += maxBatchNum
	}
	return
}

//发送有顺序的批量事件
func publishMultiEvent(sendTopic string, eventList []ReportEvent, isOrder bool) (err error) {
	kafakcfg := getKafkaConfig()
	if isOrder {
		kafakcfg = getKafkOrderConfig()
	}
	p, err := sarama.NewSyncProducer(Addresses, kafakcfg)
	if err != nil {
		log.Info("sarama.NewSyncProducer err, message=%s \n", err)
		return err
	}
	defer p.Close()
	msgs := []*sarama.ProducerMessage{}
	for _, eventData := range eventList {
		if eventData.RetryTime == "" && eventData.Num == 0 {
			eventData.Num = 1
		} else {
			eventData.Num = eventData.Num + 1
		}
		eventData.RetryTime = GetRetryTime(eventData.Num)

		jsonData, err := json.Marshal(eventData)
		if err != nil {
			log.Info("json marshal err: %+v %v", eventData, err)
			return nil
		}
		msg := &sarama.ProducerMessage{
			Topic: sendTopic,
			Value: sarama.ByteEncoder(jsonData),
			//Partition: 1,
		}
		if isOrder {
			msg.Key = sarama.ByteEncoder(eventData.Account)
		}
		msgs = append(msgs, msg)
	}

	err = p.SendMessages(msgs)
	if err != nil {
		log.Info("publish event err: %v", err)
		return err
	}
	return nil
}

/*
	记录发送失败的事件
*/
func recordFailEvent(eventData *ReportEvent) {
	_ = ReportFailEvent{
		Id:        eventData.Id,
		Type:      eventData.Type,
		Account:   eventData.Account,
		Time:      eventData.Time,
		Body:      eventData.Body,
		Topic:     eventData.Topic,
		Num:       eventData.Num,
		RetryTime: eventData.RetryTime,
	}
	//入库记录todo

}

/*
	批量记录发送失败的事件
*/
func recordBatchFailEvent(eventList []ReportEvent) {
	var events = make([]interface{}, len(eventList))
	for i, v := range eventList {
		faliEvent := ReportFailEvent{
			Id:        v.Id,
			Type:      v.Type,
			Account:   v.Account,
			Time:      v.Time,
			Body:      v.Body,
			Topic:     v.Topic,
			Num:       v.Num,
			RetryTime: v.RetryTime,
		}

		events[i] = faliEvent
	}
	//入库记录todo

}

//发送失败时，获取下次通知时间
//在通知一直不成功的情况下，总共会发起15次通知，通知频率为15s/15s/30s/3m/10m/20m/30m/30m/30m/60m/3h/3h/3h/6h/6h - 总计 24h4m

func GetRetryTime(num int32) string {
	var retryTime string
	switch num {
	case 1, 2:
		retryTime = GetNextTime("15s")
	case 3:
		retryTime = GetNextTime("30s")
	case 4:
		retryTime = GetNextTime("3m")
	case 5:
		retryTime = GetNextTime("10m")
	case 6:
		retryTime = GetNextTime("20m")
	case 7, 8, 9:
		retryTime = GetNextTime("30m")
	case 10:
		retryTime = GetNextTime("60m")
	case 11, 12, 13:
		retryTime = GetNextTime("3h")
	case 14, 15:
		retryTime = GetNextTime("6h")
		break
	default:
		retryTime = GetNextTime("15s")
	}
	return retryTime
}

//时间运算
func GetNextTime(duration string) string {
	now := time.Now()
	s, _ := time.ParseDuration(duration)
	nowAfter15Second := now.Add(s)
	t := nowAfter15Second.Format("2006-01-02 15:04:05")
	return t
}
