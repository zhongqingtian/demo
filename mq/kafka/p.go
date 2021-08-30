package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

// 基于sarama第三方库开发的kafka client

func ProductMsg(topic,ms string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		// Key:       sarama.StringEncoder(ms),
		Value:    sarama.StringEncoder(ms),
		Headers:  nil,
		Metadata: nil,
		Offset:   0,
		// Partition: 3, // 指定发送到哪个分区
		Timestamp: time.Time{},
	}

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Printf("producer closed, err:%s", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg) // 返回分区以及 偏移量。后面失败可以客户端重试
	if err != nil {
		log.Printf("send msg failed, err:%s", err)
		return
	}
	log.Printf("pid:%v offset:%v\n", pid, offset)
}
