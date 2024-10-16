package rabbit

import (
	"context"
	"log"
	"os"
	"time"

	"cursos-api/client/cursos"

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

	var err3 error
	_, err3 = Channel.QueueDeclare(
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

	for _, x := range _cursos {

		var err error
		err = Channel.PublishWithContext(ctx,
			"",             // exchange
			"cursos_queue", // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(x),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", x)
	}
}
