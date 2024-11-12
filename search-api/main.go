package main

import (
	"log"
	"search-api/client/queues"
	controllers "search-api/controllers/search"
	repo "search-api/repositories"
	services "search-api/services/search"

	"github.com/gin-gonic/gin"
)

func main() {
	config := repo.SolrConfig{
		BaseURL:    "http://localhost:8983",
		Collection: "cursos",
	}

	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "cursos_queue",
	})

	solaris := repo.NewSolr(config)

	cursosAPI := repo.NewHTTP(repo.HTTPConfig{
		Host: "localhost",
		Port: "8084",
	})

	service := services.NewService(solaris, cursosAPI, eventsQueue)

	controller := controllers.NewController(service)

	if err := eventsQueue.StartConsumer(service.HandleCursoNew); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	router := gin.Default()
	router.GET("/search", controller.Search)
	router.Run()
}
