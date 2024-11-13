package main

import (
	"log"
	"search-api/client/queues"
	controllers "search-api/controllers/search"
	repo "search-api/repositories"
	services "search-api/services/search"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := repo.SolrConfig{
		BaseURL:    "http://localhost:8983",
		Collection: "courses",
	}

	solaris := repo.NewSolr(config)

	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "cursos_queue",
	})

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/search", controller.Search)
	router.Run()
}
