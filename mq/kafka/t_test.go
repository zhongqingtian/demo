package kafka

import (
	"fmt"
	"testing"
)

func TestProductMsg(t *testing.T) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("hello %d", i)
		ProductMsg(msg)
	}
}

func TestConsumer1(t *testing.T) {
	Consumer()
}

func TestConsumer2(t *testing.T) {
	Consumer()
}