package cursos

import (
	"cursos-api/client/cursos"

	"github.com/gin-gonic/gin"
)

func GetCursos(c *gin.Context) {
	c.IndentedJSON(200, cursos.GetAllCursos())
}
