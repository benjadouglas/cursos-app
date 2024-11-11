package main

import (
	"context"
	"fmt"
	controllers "search-api/controllers/search"
	repo "search-api/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	config := repo.SolrConfig{
		BaseURL:    "http://localhost:8983",
		Collection: "cursos",
	}
	solaris := repo.NewSolr(config)

	cursos, err := solaris.Search(context.TODO(), "Id:2", 5, 10)
	if err != nil {

	}
	// Services
	fmt.Printf("%+v", cursos)

	// Hotels API
	cursosAPI := repo.NewHTTP(repo.HTTPConfig{
		Host: "cursos-api",
		Port: "8081",
	})

	service := services.NewService(solaris, cursosAPI)

	controller := controllers.NewController(service)

	router := gin.Default()
	router.GET("/search", controller.Search)
	// fmt.Printf("Courses found: %+v\n", coursesList)
}
