package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

// 初始化生产者
func InitProducer1(str string) *nsq.Producer {
	var err error
	fmt.Println("address: ", str)
	conf := nsq.NewConfig()                      // 创建一个默认配置
	producer1, err := nsq.NewProducer(str, conf) //传入IP地址，创建一个生产者
	if err != nil {
		panic(err)
	}
	return producer1
}

//发布消息
// 传参 topic 和 消息内容
func Publish1(producer1 *nsq.Producer, topic string, message string) error {
	var err error
	if producer1 != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		// 调用 生产者 推送 主题 和 消息
		err = producer1.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil %v", err)
}
