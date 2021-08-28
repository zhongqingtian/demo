package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func RunDlxP() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")
	defer ch.Close()

	args := amqp.Table{"x-dead-letter-exchange": "dlx"} // 超时转发到这个交换机，要确保这个交换机名已经存在dlx
	q, err := ch.QueueDeclare(
		"test", // 消息推给这个队列
		true,
		false,
		false,
		false,
		args) // 声明一个test队列，并设置队列的死信交换机为"dlx"

	for i := 0; i < 10; i++ {
		body := fmt.Sprintf("hello world %d", i)
		err = ch.Publish("",
			q.Name,
			false,
			false, amqp.Publishing{
			Body:       []byte(body),
			Expiration: "5000", // 设置TTL为5秒
		})
		failOnError(err, "Failed to push msg")
	}
}

func DxlConsumer() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	c, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")

	msgs, err := c.Consume(
		"dlx_queue",
		"李四",
		false,
		false,
		false,
		false,
		nil) //监听dlxQueue队列
	failOnError(err, "Failed to consum")
	ch := make(chan bool)
	defer close(ch)
	go func() {
		for d := range msgs {
			fmt.Printf("收到信息: %s\n", d.Body) // 收到消息，业务处理
			d.Ack(true)
		}
	}()

	<-ch

}
