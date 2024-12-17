package cursos

import (
	"context"
	domain "cursos-api/domain/cursos"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetCursoByID(ctx context.Context, id string) (domain.Curso, error)
	Create(ctx context.Context, curso domain.Curso) (string, error)
	Update(ctx context.Context, curso domain.Curso) error
	Delete(ctx context.Context, id string) error
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetCursoById(ctx *gin.Context) {
	// Validate ID param
	cursoID := strings.TrimSpace(ctx.Param("id"))
	// cursoID = strings.TrimPrefix(cursoID, "id:") // Add this line to remove "id:" prefix

	// Get curso by ID using the service
	curso, err := controller.service.GetCursoByID(ctx.Request.Context(), cursoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error getting curso: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, curso)
}

func (controller Controller) Create(ctx *gin.Context) {
	// Parse curso
	var curso domain.Curso
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	logrus.Printf("controller create: %v", curso)
	// Create curso
	id, err := controller.service.Create(ctx.Request.Context(), curso)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating curso: %s", err.Error()),
		})
		return
	}

	// Send ID
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (controller Controller) Update(ctx *gin.Context) {
	// Validate ID param
	id := strings.TrimSpace(ctx.Param("id"))

	// Parse curso
	var curso domain.Curso
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	// Set the ObjectID to curso
	curso.Id = id
	logrus.Printf("%v", curso)
	// Update curso
	if err := controller.service.Update(ctx.Request.Context(), curso); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error updating curso: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
		"curso":   curso,
	})
}

func (controller Controller) Delete(ctx *gin.Context) {
	// Validate ID param
	id := strings.TrimSpace(ctx.Param("id"))

	// Delete curso
	if err := controller.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error deleting curso: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}
