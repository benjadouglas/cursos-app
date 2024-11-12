package queues

import (
	"encoding/json"
	"fmt"
	"search-api/models"

	log "github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	QueueName string
}

type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbit(config RabbitConfig) Rabbit {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("error getting Rabbit connection: %v", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("error creating Rabbit channel: %v", err)
	}
	queue, err := channel.QueueDeclare(config.QueueName, false, false, false, false, nil)
	return Rabbit{
		connection: connection,
		channel:    channel,
		queue:      queue,
	}
}

func (queue Rabbit) StartConsumer(handler func(models.CursoNew)) error {
	messages, err := queue.channel.Consume(
		queue.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error registering consumer: %w", err)
	}

	go func() {
		for msg := range messages {
			log.Printf("Received message: %s", string(msg.Body))

			var cursoUpdate models.CursoNew
			if err := json.Unmarshal(msg.Body, &cursoUpdate); err != nil {
				log.Printf("error unmarshaling message: %v", err)
				continue
			}

			handler(cursoUpdate)
		}
	}()
	return nil
}

func (queue Rabbit) Publish(cursoNew models.CursoNew) error {
	bytes, err := json.Marshal(cursoNew)
	if err != nil {
		return fmt.Errorf("error marshaling Rabbit cursoNew: %w", err)
	}
	if err := queue.channel.Publish(
		"",
		queue.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		}); err != nil {
		return fmt.Errorf("error publishing to Rabbit: %w", err)
	}
	return nil
}

func (queue Rabbit) Close() {
	if err := queue.channel.Close(); err != nil {
		log.Printf("error closing Rabbit channel: %v", err)
	}
	if err := queue.connection.Close(); err != nil {
		log.Printf("error closing Rabbit connection: %v", err)
	}
}
