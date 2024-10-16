package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError( err error, msg string){
	if err != nil{
		log.Panicf("%s: %s",msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Error connecting")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Error connecting to Channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
		"cursos_queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
    )
    failOnError(err, "Failed to declare queue")

    msgs, err := ch.Consume(
        q.Name,
        "",
        true,   // auto-ack
          false,  // exclusive
      false,  // no-local
      false,  // no-wait
      nil,    // args
    )
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


