package kafka

import (
	"fmt"
	"testing"
)

func TestProductMsg(t *testing.T) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%d", i)
		ProductMsg("web_log", msg)
	}
}

// 生产者消费者模式，多个pod会百分百重复消费，不支持多节点竞争
func TestConsumer1(t *testing.T) {
	Consumer("web_log")
}

func TestConsumer2(t *testing.T) {
	Consumer("web_log")
}

// 消费者组 多个pod之间有负载作用，不重复消费，不重复

func TestConsumerGroupA1(t *testing.T) {
	ConsumerGroup("web_log", "AA")
}

func TestConsumerGroupA2(t *testing.T) {
	ConsumerGroup("web_log", "AA")
}
// 多个组，起到订阅发布一对多作用
func TestConsumerGroupB1(t *testing.T) {
	ConsumerGroup("web_log", "BB")
}

func TestConsumerGroupB2(t *testing.T) {
	ConsumerGroup("web_log", "BB")
}

func TestConsumerGroupB3(t *testing.T) {
	ConsumerGroup("web_log", "BB")
}
