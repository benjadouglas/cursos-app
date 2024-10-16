package main

import (
	"log"

	"cursos-api/db"
	"cursos-api/rabbit"
	"cursos-api/router"

	"github.com/gin-gonic/gin"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// mongo setup
	defer db.Close()

	// rabbit setup
	rabbit.Connect()
	defer rabbit.Close()
	rabbit.Migrate()

	// routing setup
	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
}
