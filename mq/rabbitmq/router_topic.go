package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func TopicP(exchangeName, routerKey, msg string) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		exchangeName, // exchange
		routerKey,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" routerKey=%s [x] Sent %s", routerKey, msg)
}

func RunTopicConsumer(exchangeName, queueName string, routerKeys []string) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, s := range routerKeys {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, exchangeName, s)
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,     // queue
		"zhangsan", // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	defer close(forever)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
