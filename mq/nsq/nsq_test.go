package nsq

import (
	"fmt"
	"testing"
)

func TestSender(t *testing.T) {
	m := "[%d]hello work"
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf(m, i)
		Sender("test", msg)
	}
}

// 相同channel 组成一个消费组，竞争获取队列消息
func TestReceiver1(t *testing.T) {
	Receiver("test", "ch1")
}

func TestReceiver2(t *testing.T) {
	Receiver("test", "ch1")
}

// 不同channel 起到订阅发布 一对多 广播
func TestReceiver3(t *testing.T) {
	Receiver("test", "ch2")
}
