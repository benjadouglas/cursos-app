package rabbit

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"cursos-api/client/cursos"
	"cursos-api/model"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Connection *amqp.Connection
var Channel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró el archivo .env")
	}
	uri := os.Getenv("RABBITMQ_URI")
	if uri == "" {
		log.Fatal("Debes establecer la variable de entorno 'RABBITMQ_URI'")
	}

	var err1 error
	Connection, err1 = amqp.Dial(uri)
	failOnError(err1, "Failed to connect to RabbitMQ")

	var err2 error
	Channel, err2 = Connection.Channel()
	failOnError(err2, "Failed to open a channel")

	_, err3 := Channel.QueueDeclare(
		"cursos_queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err3, "Failed to declare a queue")
}

func Close() {
	if err := Channel.Close(); err != nil {
		log.Printf("Error al cerrar el canal: %v", err)
	}
	if err := Connection.Close(); err != nil {
		log.Printf("Error al cerrar la conexión: %v", err)
	}
}

func Migrate() {
	_cursos, err := cursos.GetCursos()
	failOnError(err, "Failed fetching cursos")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, curso := range _cursos {
		jsonData, err1 := json.Marshal(curso)
		if err1 != nil {
			log.Printf("Error marshaling curso to JSON: %v", err1)
			continue
		}

		err2 := Channel.PublishWithContext(ctx,
			"",             // exchange
			"cursos_queue", // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(jsonData),
			})
		failOnError(err2, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", jsonData)
	}
}

func Publish(curso model.Curso) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err1 := json.Marshal(curso)
	if err1 != nil {
		failOnError(err1, "Error marshaling curso to Json")
	}

	err2 := Channel.PublishWithContext(ctx,
		"",             // exchange
		"cursos_queue", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(jsonData),
		})
	failOnError(err2, "Failed to publish message")
	log.Printf(" [x] Sent %s\n", jsonData)
}
