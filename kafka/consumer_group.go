package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Kafka struct {
	brokers []string
	topics  []string
	//OffsetNewest int64 = -1
	//OffsetOldest int64 = -2
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
}

func NewKafka() *Kafka {
	return &Kafka{
		brokers: brokers,
		topics: []string{
			topics,
		},
		group:             group,
		channelBufferSize: 2,
		ready:             make(chan bool),
		version:           "1.1.1",
	}
}

var brokers = []string{"127.0.0.1:9092"}
var topics = "test"
var group = "1"

func (p *Kafka) Init() func() {
	log.Infoln("kafka init...")

	version, err := sarama.ParseKafkaVersion(p.version)
	if err != nil {
		log.Fatalf("Error parsing Kafka version: %v", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.channelBufferSize                         // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(p.brokers, p.group, config) // 创建消费者组的客户端
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			//util.HandlePanic("client.Consume panic", log.StandardLogger())
		}()
		for {
			if err := client.Consume(ctx, p.topics, p); err != nil { // p 要实现  ConsumerGroupHandler 接口
				log.Fatalf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			p.ready = make(chan bool)
		}
	}()
	<-p.ready
	log.Infoln("Sarama consumer up and running!...")
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Info("kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Errorf("Error closing client: %v", err)
		}
	}
}

// ---------------------------------------------------------
// 实现消费组的三个接口

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// 具体消费消息
	for message := range claim.Messages() {
		msg := string(message.Value)
		log.Infof("msg: %s", msg)
		time.Sleep(time.Second)
		//run.Run(msg)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

// --------------------------------------------------------
func ConsumerGroup() {
	k := NewKafka()
	f := k.Init()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	f()
}
