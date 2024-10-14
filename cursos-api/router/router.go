package router

import (
	"cursos-api/controller/cursos"

	"github.com/gin-gonic/gin"
)

func createCurso(c *gin.Context) {
	println("Creaste un curso")
}

func MapUrls(engine *gin.Engine) {
	engine.GET("/cursos", cursos.GetCursos)
}
