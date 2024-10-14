package router

import "github.com/gin-gonic/gin"

func createCurso(c *gin.Context) {
	println("Creaste un curso")
}

func MapUrls(engine *gin.Engine) {
	engine.POST("/curso", createCurso)
	engine.GET("/curso/:id", getCursoById)
}
