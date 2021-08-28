package rabbitmq

import (
	"fmt"
	"testing"
	"time"
)

// 默认queue 多个pod 可以消费同一个queue 轮询
func TestRunProducer(t *testing.T) {
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("Hello World! %d", i)
		RunProducer(msg)
		// time.Sleep(1*time.Second)
	}
}

func TestRunConsumer1(t *testing.T) {
	ch := make(chan bool)
	RunConsumer(ch)

	time.Sleep(2 * time.Minute)
	ch <- true
}

func TestRunConsumer2(t *testing.T) {
	ch := make(chan bool)
	RunConsumer(ch)

	time.Sleep(2 * time.Minute)
	ch <- true
}

// fanout
func TestRunExchangeP(t *testing.T) {
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("Hello World! %d", i)
		RunExchangeP(msg)
		// time.Sleep(1*time.Second)
	}
}

func TestRunExchangeConsumer(t *testing.T) {
	forever := make(chan bool)
	queueName := "q1"
	RunExchangeConsumer(forever, queueName)
	time.Sleep(2 * time.Minute)
	forever <- true
}

func TestRunExchangeConsumer2(t *testing.T) {
	forever := make(chan bool)
	queueName := "q2" // 不能在相同交换机里面 重复注册队列名
	RunExchangeConsumer(forever, queueName)
	time.Sleep(2 * time.Minute)
	forever <- true
}

// router key direct
func TestRouterP(t *testing.T) { // 一个消息属于多个 routerKey，要发多次
	routerKey := "warn"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		RouterP(routerKey, msg)
		// time.Sleep(1*time.Second)
	}
	routerKey = "err"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		RouterP(routerKey, msg)
		// time.Sleep(1*time.Second)
	}
	routerKey = "info"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		RouterP(routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

func TestRunRouterConsumer1(t *testing.T) {
	routerKeys := []string{"warn", "err"}
	RunRouterConsumer(routerKeys)
}

func TestRunRouterConsumer2(t *testing.T) {
	routerKeys := []string{"warn"}
	RunRouterConsumer(routerKeys)
}

func TestRunRouterConsumer3(t *testing.T) {
	routerKeys := []string{"warn"}
	RunRouterConsumer(routerKeys)
}

// 四 topic

func TestTopicP1(t *testing.T) {
	routerKey := "quick.orange.rabbit"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		TopicP(routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

func TestTopicP2(t *testing.T) {
	routerKey := "lazy.orange.elephant"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%s=! %d", routerKey, i)
		TopicP(routerKey, msg)
		// time.Sleep(1*time.Second)
	}
}

func TestRunTopicConsumer(t *testing.T) {
	routerKeys := []string{"*.orange.*"}
	RunTopicConsumer(routerKeys)
}

// 死信队列
func TestRunDlxP(t *testing.T) {
	RunDlxP()
}

func TestDxlConsumer(t *testing.T) {
	DxlConsumer()
}