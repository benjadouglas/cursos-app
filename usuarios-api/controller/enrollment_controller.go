package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"usuarios-api/model"

	"github.com/sirupsen/logrus"
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

	logrus.Printf("Json data %v", enrollment)
	// Verificar si el usuario existe
	var user model.Usuario
	if err := db.First(&user, enrollment.Id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verificar si ya existe la inscripción
	var existingEnrollment model.Enrollment
	if err := db.Where("id = ? AND id_cursos = ?", enrollment.Id, enrollment.Id_cursos).
		First(&existingEnrollment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enrollment already exists"})
		return
	}

	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create enrollment"})
		return
	}

	// Count total enrollments for this course
	var count int64
	if err := db.Model(&model.Enrollment{}).Where("id_cursos = ?", enrollment.Id_cursos).Count(&count).Error; err != nil {
		log.Printf("Error counting enrollments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count enrollments"})
		return
	}

	// Update course capacity in cursos-api
	updateURL := fmt.Sprintf("http://localhost:8084/cursos/%s", enrollment.Id_cursos)
	updatePayload := map[string]interface{}{
		"Capacidad": count,
	}
	jsonPayload, _ := json.Marshal(updatePayload)

	req, err := http.NewRequest("PUT", updateURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating request to update course capacity: %v", err)
	} else {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Printf("Error updating course capacity: %v", err)
		} else {
			defer response.Body.Close()
		}
	}

	c.JSON(http.StatusCreated, enrollment)
}

// GetEnrollmentsByUserID obtiene todas las inscripciones de un usuario
func GetEnrollmentsByUserID(c *gin.Context) {
	userID := strings.TrimSpace(c.Param("id"))
	logrus.Printf("%s", userID)
	userID = strings.TrimPrefix(userID, "id:") // Add this line to remove "id:" prefix
	logrus.Printf("%v", userID)

	var enrollments []model.Enrollment
	if err := db.Where("id = ?", userID).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve enrollments"})
		return
	}

	courses := make([]interface{}, 0)
	for _, enrollment := range enrollments {
		resp, err := http.Get("http://localhost:8084/cursos/" + enrollment.Id_cursos)
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

	// Get the enrollment first to have the course ID
	var enrollment model.Enrollment
	if err := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	// Delete the enrollment
	if err := db.Delete(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete enrollment"})
		return
	}

	// Count remaining enrollments for this course
	var count int64
	if err := db.Model(&model.Enrollment{}).Where("id_cursos = ?", enrollment.Id_cursos).Count(&count).Error; err != nil {
		log.Printf("Error counting enrollments: %v", err)
	} else {
		// Update course capacity in cursos-api
		updateURL := fmt.Sprintf("http://cursos-api:8084/cursos/%s", enrollment.Id_cursos)
		updatePayload := map[string]interface{}{
			"Capacidad": count,
		}
		jsonPayload, _ := json.Marshal(updatePayload)

		req, err := http.NewRequest("PUT", updateURL, bytes.NewBuffer(jsonPayload))
		if err != nil {
			log.Printf("Error creating request to update course capacity: %v", err)
		} else {
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				log.Printf("Error updating course capacity: %v", err)
			} else {
				defer response.Body.Close()
			}
		}
	}

	c.Status(http.StatusNoContent)
}
