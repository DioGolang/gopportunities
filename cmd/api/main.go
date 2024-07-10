package main

import (
	"github.com/devsvasconcelos/gopportunities.git/internal/infrastructure/database"
	"github.com/devsvasconcelos/gopportunities.git/internal/interfaces"
	handle "github.com/devsvasconcelos/gopportunities.git/internal/interfaces/handler"
	"github.com/devsvasconcelos/gopportunities.git/internal/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	jobRepo := interfaces.NewJobRepository(db)
	jobUseCase := usecases.NewJobUseCase(jobRepo)
	jobHandler := handle.NewJobHandler(jobUseCase)

	router := gin.Default()

	router.POST("/jobs", jobHandler.CreateJob)
	router.GET("/jobs", jobHandler.GetJobs)

	router.Run(":8080")
}
