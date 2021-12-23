package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	addr := []string{"127.0.0.1:32768"}

	producer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "hello",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			break
		}
		msg.Value = sarama.ByteEncoder(value)
		fmt.Println(value)

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send message Fail")
		}
		fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}

//同步消息模式
func syncProducer(address []string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//  // 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Timeout = 5 * time.Second

	// 使用给定代理地址和配置创建一个同步生产者
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	topic := "test"
	srcValue := "sync: this is a message. index=%d"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		//SendMessage：该方法是生产者生产给定的消息
		//生产成功的时候返回该消息的分区和所在的偏移量
		//生产失败的时候返回error
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2 * time.Second)
	}
}

//异步消费者(Goroutines)：用不同的goroutine异步读取Successes和Errors channel
func asyncProducer1(address []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	//Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	var enqueued, successes, errors int
	wg.Add(2) //2 goroutine

	// 发送成功message计数
	go func() {
		defer wg.Done()
		for range p.Successes() {
			successes++
		}
	}()

	// 发送失败计数
	go func() {
		defer wg.Done()
		for err := range p.Errors() {
			log.Printf("%v 发送失败，err：%s\n", err.Msg, err.Err)
			errors++
		}
	}()

	// 循环发送信息
	asrcValue := "async-goroutine: this is a message. index=%d"
	var i int
Loop:
	for {
		i++
		value := fmt.Sprintf(asrcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: "test",                    // 主题
			Value: sarama.ByteEncoder(value), // 值
		}
		select {
		case p.Input() <- msg: // 发送消息
			enqueued++
			fmt.Fprintln(os.Stdout, value)
		case <-signals: // 中断信号
			p.AsyncClose()
			break Loop
		}
		time.Sleep(2 * time.Second)
	}
	wg.Wait()

	fmt.Fprintf(os.Stdout, "发送数=%d，发送成功数=%d，发送失败数=%d \n", enqueued, successes, errors)

}
