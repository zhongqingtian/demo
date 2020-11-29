package nsq

import (
	"github.com/nsqio/go-nsq"
)

// publisher
type Producer interface {
	Publish(topic string, data []byte) error
}

type NSQProducer struct {
	P *nsq.Producer
}

func NewNSQProducer(addr string) (Producer, error) {
	var err error
	np := new(NSQProducer)
	np.P, err = nsq.NewProducer(addr, nsq.NewConfig())
	return np, err
}

func (np *NSQProducer) Publish(topic string, data []byte) error {
	//bts, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}
	return np.P.Publish(topic, data)
}
