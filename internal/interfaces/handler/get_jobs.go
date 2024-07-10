package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *JobHandler) GetJobs(c *gin.Context) {
	jobs, err := h.jobUseCase.GetJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}
