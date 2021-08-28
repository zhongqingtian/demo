package nsq

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

const (
	NUM_CURRENT_HANDLER = 100
)

type MessageHandler interface {
	HandleMessage(msgID string, msgBody []byte) error
}

type NSQConsumer struct {
	C *nsq.Consumer

	addrs   []string
	topic   string
	channel string

	handler MessageHandler
}

func NewNSQConsumer(cfg *NSQConfig, topic, channel string) (*NSQConsumer, error) {
	var err error
	nc := new(NSQConsumer)
	nsqCfg := nsq.NewConfig()
	nsqCfg.MaxInFlight = 50
	c, err := nsq.NewConsumer(topic, channel, nsqCfg)
	if err != nil {
		return nil, err
	}
	nc.C = c
	nc.addrs = cfg.LookupAddrs
	nc.topic = topic
	nc.channel = channel
	return nc, err
}

func (nc *NSQConsumer) HandleMessage(msg *nsq.Message) error {
	return nc.handler.HandleMessage(fmt.Sprintf("%x", msg.ID), msg.Body)
}

func (nc *NSQConsumer) ListenMessage(handler MessageHandler) error {
	nc.handler = handler
	nc.C.AddConcurrentHandlers(nc, NUM_CURRENT_HANDLER)
	return nc.C.ConnectToNSQLookupds(nc.addrs)
}
