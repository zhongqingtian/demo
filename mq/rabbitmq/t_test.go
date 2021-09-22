package rabbitmq

import (
	"fmt"
	"testing"
)

// 默认queue 多个pod 可以消费同一个queue 轮询,不同消费者名 组成消费组
// 消费者 重启后会把未确认的消息退给其他同组消费者，或者自己重启完，退给自己
func TestRunProducer(t *testing.T) {
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("%d", i)
		RunProducer("hello", msg)
	}
}

func TestRunConsumer1(t *testing.T) {
	RunConsumer("hello", "consumer")
}

func TestRunConsumer2(t *testing.T) {
	RunConsumer("hello", "consumer")
}

func TestRunConsumer3(t *testing.T) {
	RunConsumer("hello", "consumer3")
}

// ------------------

// fanout
func TestRunExchangeP(t *testing.T) {
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("Hello World! %d", i)
		RunExchangeP("logs", msg)
		// time.Sleep(1*time.Second)
	}
}

// 多个pod 可以轮询消费 同一个队列，竞争关系
func TestRunExchangeConsumer1(t *testing.T) {
	RunExchangeConsumer("logs", "q1", "consumer")
}

func TestRunExchangeConsumer2(t *testing.T) {
	RunExchangeConsumer("logs", "q1", "consumer")
}

// router key direct
func TestRouterP(t *testing.T) { // 一个消息属于多个 routerKey，要发多次
	routerKey := "warn"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("[%d] %s=!", i, routerKey)
		RouterP("logs_direct", routerKey, msg)
		// time.Sleep(1*time.Second)
	}
	routerKey = "err"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("[%d] %s=!", i, routerKey)
		RouterP("logs_direct", routerKey, msg)
		// time.Sleep(1*time.Second)
	}
	routerKey = "info"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("[%d] %s=!", i, routerKey)
		RouterP("logs_direct", routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

// 相同服务 多个节点（pod）会竞争一个队列消息
func TestRunRouterConsumer1(t *testing.T) {
	routerKeys := []string{"warn", "err"}
	RunRouterConsumer("logs_direct", "q111", routerKeys)
}

func TestRunRouterConsumer11(t *testing.T) {
	routerKeys := []string{"warn", "err"}
	RunRouterConsumer("logs_direct", "q111", routerKeys)
}

// 不同队列订阅发布 起到一对多广播作用
func TestRunRouterConsumer2(t *testing.T) {
	routerKeys := []string{"warn"}
	RunRouterConsumer("logs_direct", "q222", routerKeys)
}

func TestRunRouterConsumer3(t *testing.T) {
	routerKeys := []string{"info", "err"}
	RunRouterConsumer("logs_direct", "q333", routerKeys)
}

// 四 topic
func TestTopicP1(t *testing.T) {
	routerKey := "quick.orange.rabbit"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		TopicP("logs_direct", routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

func TestTopicP2(t *testing.T) {
	routerKey := "lazy.orange.elephant"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("[%d] %s=! ", i, routerKey)
		TopicP("logs_topic", routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

func TestRunTopicConsumer1(t *testing.T) {
	routerKeys := []string{"*.orange.*"}
	RunTopicConsumer("logs_topic","queue444",routerKeys)
}

func TestRunTopicConsumer2(t *testing.T) {
	routerKeys := []string{"*.orange.*"}
	RunTopicConsumer("logs_topic","queue444",routerKeys)
}

// 死信队列
func TestRunDlxP(t *testing.T) {
	m := "hello [%d]"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf(m, i)
		ProducerDlx("long_abc", msg)
	}
}

func TestConsumer(t *testing.T) {
	Consumer("long_abc", "dlx_exchange")
}

func TestConsumerDlx(t *testing.T) {
	ConsumerDlx("dlx_exchange", "dlx_queue")
}
