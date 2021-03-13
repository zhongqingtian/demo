package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func Consumer() {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		panic(err)
	}
	//     //Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions("test") // 一个topic 下有多个分区
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest) // 每个分区从原始位置开始消费
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(pc sarama.PartitionConsumer) { // 多个分区消费者同时消费
			defer wg.Done()
			// Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
