package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func RunProducer(msg string) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// body := "Hello World! 1"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

func RunConsumer(queueName, consumerName string) {
	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	/*err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)*/
	msgs, err := ch.Consume(
		q.Name,
		consumerName,
		false,
		false,
		false,
		false, nil)
	failOnError(err, "Failed to declare a queue")
	go func() {
		//	i := 0
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// time.Sleep(3*time.Second)
			d.Ack(true) // 重启后会把未确认的消息退给其他同组消费者，或者自己重启完，退给自己
			/*if i/2==1 {
				d.Ack(true)
			} else {
				d.Ack(false)
			}*/
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
