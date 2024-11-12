package main

import (
	log "github.com/sirupsen/logrus"

	controllers "cursos-api/controller/cursos"
	queues "cursos-api/rabbit"
	repositories "cursos-api/repo/cursos"
	services "cursos-api/services/cursos"

	"github.com/gin-gonic/gin"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// mongo setup
	//  MONGODB_URI="mongodb://root:root@localhost:27017"
	mainRepo := repositories.NewMongo(repositories.MongoConfig{
		Host:       "localhost",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "cursos-api",
		Collection: "cursos",
	})

	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "cursos_queue",
	})

	service := services.NewService(mainRepo, eventsQueue)

	controller := controllers.NewController(service)

	router := gin.Default()
	router.GET("/cursos/:id", controller.GetCursoById)
	router.POST("/cursos", controller.Create)
	router.PUT("/cursos/:id", controller.Update)
	router.DELETE("/cursos/:id", controller.Delete)
	if err := router.Run(":8084"); err != nil {
		log.Fatalf("error running application: %v", err)
	}

}
