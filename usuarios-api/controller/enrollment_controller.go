package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usuarios-api/model"
)

// CreateEnrollment crea una nueva inscripción
func CreateEnrollment(c *gin.Context) {
	var enrollment model.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

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
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var enrollments []model.Enrollment
	if err := db.Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve enrollments"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
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
