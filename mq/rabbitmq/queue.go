package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
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

func RunConsumer(forever chan bool) {
	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	/*err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)*/
	msgs, err := ch.Consume(
		q.Name,
		"test1",
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
			// d.Ack(true)
			/*if i/2==1 {
				d.Ack(true)
			} else {
				d.Ack(false)
			}*/
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
