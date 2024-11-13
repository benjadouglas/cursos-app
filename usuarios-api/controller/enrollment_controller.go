package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"usuarios-api/model"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// CreateEnrollment crea una nueva inscripción
func CreateEnrollment(c *gin.Context) {
	var enrollment model.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Json data %v", enrollment)
	// Verificar si el usuario existe
	var user model.Usuario
	if err := db.First(&user, enrollment.Id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verificar si ya existe la inscripción
	var existingEnrollment model.Enrollment
	if err := db.Where("user_id = ? AND course_id = ?", enrollment.Id, enrollment.Id_cursos).
		First(&existingEnrollment).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Enrollment already exists"})
		return
	}

	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create enrollment"})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

// GetEnrollmentsByUserID obtiene todas las inscripciones de un usuario
func GetEnrollmentsByUserID(c *gin.Context) {
	userID := strings.TrimSpace(c.Param("id"))
	userID = strings.TrimPrefix(userID, "id:") // Add this line to remove "id:" prefix
	log.Printf("%v", userID)

	var enrollments []model.Enrollment
	if err := db.Where("id = ?", userID).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve enrollments"})
		return
	}

	courses := make([]interface{}, 0)
	for _, enrollment := range enrollments {
		resp, err := http.Get("http://localhost:8084/cursos/id:" + enrollment.Id_cursos)
		if err != nil {
			log.Printf("Error getting course: %v", err)
			continue
		}
		defer resp.Body.Close()

		var course interface{}
		if err := json.NewDecoder(resp.Body).Decode(&course); err != nil {
			log.Printf("Error decoding response: %v", err)
			continue
		}
		courses = append(courses, course)
	}
	log.Printf("cursos:\n %v", courses)
	c.JSON(http.StatusOK, courses)
}

// DeleteEnrollment elimina una inscripción
func DeleteEnrollment(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	courseID, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	result := db.Where("user_id = ? AND course_id = ?", userID, courseID).Delete(&model.Enrollment{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete enrollment"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
