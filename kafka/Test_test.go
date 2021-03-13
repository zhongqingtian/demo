package kafka

import "testing"

func TestProducer(t *testing.T) {
	// 	Producer()

	// 	syncProducer([]string{"127.0.0.1:9092"})
	asyncProducer1([]string{"127.0.0.1:9092"})
}

func TestConsumer(t *testing.T) {
	Consumer()
}

func TestConsumerGroup(t *testing.T) {
	ConsumerGroup()
}
