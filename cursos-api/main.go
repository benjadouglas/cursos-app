package main

import (
	"cursos-api/db"
	"cursos-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	engine := gin.New()

	router.MapUrls(engine)
	engine.Run(":8080")
}
