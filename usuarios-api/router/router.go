package router

import (
	users "usuarios-api/controller"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.POST("/api/login", users.Login)
	engine.POST("/api/users", users.CreateUser)
	engine.POST("/api/validate", users.AuthMiddleware(), users.IsAdmin)
	protected := engine.Group("/api")
	protected.Use(users.AuthMiddleware())
	{
		protected.GET("/users", users.GetAllUsers)
		protected.GET("/users/:id", users.GetUserByID)
		protected.PUT("/users/:id", users.UpdateUser)
		protected.DELETE("/users/:id", users.DeleteUser)

		protected.POST("/enrollments", users.CreateEnrollment)
		protected.GET("/enrollments/user/:id", users.GetEnrollmentsByUserID)
		// protected.PUT("/enrollments/:id", users.UpdateEnrollment)
		protected.DELETE("/enrollments/:userId/:courseId", users.DeleteEnrollment)
	}
}
