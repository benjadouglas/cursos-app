package main

import (
	"context"
	"log"
	"time"

	"cursos-api/db"
	"cursos-api/rabbit"
	"cursos-api/router"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	defer db.Close()
	engine := gin.New()
	router.MapUrls(engine)

	rabbit.Connect()
	defer rabbit.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err := rabbit.Channel.PublishWithContext(ctx,
		"",             // exchange
		"cursos_queue", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	engine.Run(":8080")
}
