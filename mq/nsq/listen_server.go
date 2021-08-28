package nsq

import "github.com/sirupsen/logrus"

type ListenServer struct {
	cfg     *NSQConfig
	listens map[string]map[string]MessageHandler // topic -> [ channel -> handle ]
}

func NewListenServer(cfg *NSQConfig) *ListenServer {
	logrus.WithField("nsq_config", cfg).Info("nsq config show")
	ls := new(ListenServer)
	ls.cfg = cfg
	return ls
}

func (ls *ListenServer) AddListener(topic, channel string, l MessageHandler) {
	if ls.listens == nil {
		ls.listens = make(map[string]map[string]MessageHandler)
	}
	if _, ok := ls.listens[topic]; ok {
		ls.listens[topic][channel] = l
	} else {
		ls.listens[topic] = map[string]MessageHandler{
			channel: l,
		}
	}
}

func (ls *ListenServer) Listen() error {
	for t, cm := range ls.listens {
		for c, l := range cm {
			cus, err := NewNSQConsumer(ls.cfg, t, c)
			if err != nil {
				logrus.WithError(err).Fatal("create nsqConsumer fail")
				return err
			}
			if err = cus.ListenMessage(l); err != nil {
				return err
			}
		}
	}
	return nil
}
