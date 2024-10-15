package router

import (
	"cursos-api/controller/cursos"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.GET("/cursos", cursos.GetCursos)
	engine.GET("/cursos/:id", cursos.GetCursoById)
}
