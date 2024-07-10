package handle

import (
	"github.com/devsvasconcelos/gopportunities.git/pkg/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *JobHandler) CreateJob(c *gin.Context) {
	var job entities.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.jobUseCase.CreateJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
