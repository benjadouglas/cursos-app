package router

import (
	"cursos-api/controller/cursos"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine, controller cursos.Controller) {

	// engine.GET("/cursos", controller.GetCursos)
	engine.GET("/cursos/:id", controller.GetCursoById)
	engine.PUT("/cursos/:id", controller.Update)
	engine.DELETE("/cursos/:id", controller.Delete)
	engine.POST("/cursos", controller.Create)
}
