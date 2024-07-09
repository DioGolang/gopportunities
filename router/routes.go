package router

import (
	"github.com/devsvasconcelos/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
	}

}
