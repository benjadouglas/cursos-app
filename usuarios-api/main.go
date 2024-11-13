package main

import (
	"log"
	"os"
	"time"
	database "usuarios-api/DBCONFIG" // Import your database package
	users "usuarios-api/controller"
	"usuarios-api/model"
	"usuarios-api/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Verificar JWT_SECRET
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is required in .env file")
	}

	// Use your database connection function
	db := database.ConnectDB()

	// Auto migrate the database
	db.AutoMigrate(&model.Usuario{})

	// Set the database for the controllers
	users.SetDB(db)

	// Initialize Gin router
	r := gin.Default()

	// Setup CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Map all routes
	router.MapUrls(r)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}
	r.Run(":" + port)
}
