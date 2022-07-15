package queue

import (
	"github.com/nsqio/go-nsq"
)

// InitConsumer 初始化一个消费者
func InitConsumer(addr, topic, channel string, handlers ...nsq.Handler) error {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	// 设置消息处理函数
	for _, handler := range handlers {
		consumer.AddHandler(handler)
	}
	if err := consumer.ConnectToNSQD(addr); err != nil {
		return err
	}
	<-consumer.StopChan
	return nil
}

// InitProducer 初始化一个消费者
func InitProducer(addr string) (*nsq.Producer, error) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
