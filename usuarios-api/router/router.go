package router

import (
	users "usuarios-api/controller"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.POST("/api/login", users.Login)
	engine.POST("/api/users", users.CreateUser)

	protected := engine.Group("/api")
	protected.Use(users.AuthMiddleware())
	{
		protected.GET("/users", users.GetAllUsers)
		protected.GET("/users/:id", users.GetUserByID)
		protected.PUT("/users/:id", users.UpdateUser)
		protected.DELETE("/users/:id", users.DeleteUser)

		protected.POST("/enrollments", users.CreateEnrollment)
		protected.GET("/enrollments/user/:id", users.GetEnrollmentsByUserID)
		protected.DELETE("/enrollments/:userId/:courseId", users.DeleteEnrollment)
	}
}
