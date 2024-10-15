package cursos

import (
	"cursos-api/client/cursos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCursos(c *gin.Context) {
	cursos, err := cursos.GetCursos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cursos"})
		return
	}
	c.IndentedJSON(http.StatusOK, cursos)
}

func GetCursoById(c *gin.Context) {
	id := c.Param("id")
	curso, err := cursos.GetCursoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, curso)
}
