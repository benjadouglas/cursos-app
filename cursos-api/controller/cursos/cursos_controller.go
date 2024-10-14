package cursos

import "github.com/gin-gonic/gin"

func GetCursos(c *gin.Context) {
	c.IndentedJSON(200, cursos.getAllCursos)
}
