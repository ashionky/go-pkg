package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"go-pkg/pkg/log"
	"sync"

)
type kafkaHandler interface {
	HandleKafkaMsg(message *ReportEvent) error
}

func NewConsumerHandler(topic string, groupID string, ctx1 context.Context, w *sync.WaitGroup, handler sarama.ConsumerGroupHandler) {
	// 当前消费者退出时，通知主线程
	defer w.Done()
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.ChannelBufferSize = 1
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Fetch.Max = 11000000

	topics := []string{topic}
	ctx, cancel := context.WithCancel(context.Background())
	consumerGroup, err := sarama.NewConsumerGroup(Addresses, groupID, config)
	if err != nil {
		log.Info("Create consumerGroup error, %v", err)
	}
	// 子协程的WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			// 为了重试
			if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
				log.Info("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	// 记录错误
	go func() {
		for err := range consumerGroup.Errors() {
			log.Error("Kafka consumerGroup error, %v", err)
		}
	}()
	// 监听主协程的上下文
	select {
	case <-ctx1.Done():
		log.Error("terminating: context cancelled")
	}
	// 有中断信号，就关闭当前的上下文，通知消费者停止拉取
	cancel()
	// 等待消费者处理完已拉取下来的信息
	wg.Wait()
	if err = consumerGroup.Close(); err != nil {
		log.Error("Error closing client: %v", err)
	}
}
