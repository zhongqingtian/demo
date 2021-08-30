package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ProducerDlx(exchangeName, msg string) {
	var (
		conn *amqp.Connection
		err  error
		ch   *amqp.Channel
	)
	if conn, err = amqp.Dial("amqp://admin:admin@127.0.0.1:5672/"); err != nil {
		log.Printf("amqp.Dial err :%s", err)
		return
	}
	defer conn.Close()

	if ch, err = conn.Channel(); err != nil {
		log.Printf("conn.Channel err: %s", err)
		return
	}
	defer ch.Close()

	//func (ch *Channel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table) error
	//声明交换器
	if err = ch.ExchangeDeclare(
		exchangeName,        // Exchange names
		amqp.ExchangeDirect, //"direct", "fanout", "topic" and "headers"
		true,
		false, //Durable and Non-Auto-Deleted exchanges会一直保留
		false,
		false,
		nil,
	); err != nil {
		log.Printf("ch.ExchangeDeclare err: %s", err)
		return
	}

	//func (ch *Channel) Publish(exchange, key string, mandatory, immediate bool, msg Publishing) error
	//发送消息
	if err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:      amqp.Table{},
			ContentType:  "text/plain",
			Body:         []byte(msg),
			DeliveryMode: amqp.Persistent, //需要做持久化保留
			Priority:     0,
			Expiration:   "9000",
		},
	); err != nil {
		fmt.Println("ch.Publish err: ", err)
		return
	}
}

func Consumer(exchangeName, dlxExchangeName string) {
	var (
		conn  *amqp.Connection
		err   error
		ch    *amqp.Channel
		queue amqp.Queue
	)
	if conn, err = amqp.Dial("amqp://admin:admin@127.0.0.1:5672/"); err != nil {
		fmt.Println("amqp.Dial err :", err)
		return
	}
	defer conn.Close()

	if ch, err = conn.Channel(); err != nil {
		fmt.Println("conn.Channel err: ", err)
		return
	}

	defer ch.Close()

	//func (ch *Channel) Qos(prefetchCount, prefetchSize int, global bool) error
	//设置未确认的最大消息数
	if err = ch.Qos(1, 0, false); err != nil {
		fmt.Println("ch.Qos err: ", err)
		return
	}

	// dlxExchangeName = "dlx_exchange"

	//声明交换器
	if err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Printf("ch.ExchangeDeclare err: %s", err)
		return
	}

	argsQue := make(map[string]interface{})
	//添加死信队列交换器属性
	argsQue["x-dead-letter-exchange"] = dlxExchangeName
	//指定死信队列的路由key，不指定使用队列路由键
	//argsQue["x-dead-letter-routing-key"] = "zhe_mess"
	//添加过期时间
	// argsQue["x-message-ttl"] = 6000 //单位毫秒
	//声明队列
	queue, err = ch.QueueDeclare("zhe_123", true, false, false, false, argsQue)
	if err != nil {
		fmt.Println("ch.QueueDeclare err :", err)
		return
	}

	//绑定交换器/队列和key
	//func (ch *Channel) QueueBind(name, key, exchange string, noWait bool, args Table) error
	if err = ch.QueueBind(queue.Name, "", exchangeName, false, nil); err != nil {
		log.Printf("ch.QueueBind err: %s", err)
		return
	}
	//开启推模式消费
	//func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table) (<-chan Delivery, error)
	delvers, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("ch.Consume err: %s", err)
	}

	for delver := range delvers {
		log.Println(string(delver.Body))
		delver.Ack(true)
		time.Sleep(5 * time.Second)
	}

}
func ConsumerDlx(dlxExchangeName, dlxQueueName string) {

	var (
		conn  *amqp.Connection
		ch    *amqp.Channel
		queue amqp.Queue
		err   error
	)

	//链接rbmq
	if conn, err = amqp.Dial("amqp://admin:admin@127.0.0.1:5672/"); err != nil {
		fmt.Println("amqp.Dial err: ", err)
		return
	}

	//声明信道
	if ch, err = conn.Channel(); err != nil {
		fmt.Println("conn.Channel err: ", err)
		return
	}

	//声明交换机
	if err = ch.ExchangeDeclare(
		dlxExchangeName,
		amqp.ExchangeFanout, //交换机模式fanout
		true,                //持久化
		false,               //自动删除
		false,               //是否是内置交互器,(只能通过交换器将消息路由到此交互器，不能通过客户端发送消息
		false,
		nil,
	); err != nil {
		log.Printf("ch.ExchangeDeclare: %s", err)
		return
	}

	//声明队列
	if queue, err = ch.QueueDeclare(
		dlxQueueName, //队列名称
		true,         //是否是持久化
		false,        //是否不需要确认，自动删除消息
		false,        //是否是排他队列
		false,        //是否等待服务器返回ok
		nil,
	); err != nil {
		log.Printf("ch.QueueDeclare err: %s", err)
		return
	}

	//将交换器和队列/路由key绑定
	if err = ch.QueueBind(queue.Name, "", dlxExchangeName, false, nil); err != nil {
		log.Printf("ch.QueueBind err: %s", err)
		return
	}

	//开启推模式消费
	delvers, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	for delver := range delvers {
		log.Println(string(delver.Body))
		delver.Ack(true)
		time.Sleep(1 / 2 * time.Second)
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Println("terminating: via signal")
	}

}
