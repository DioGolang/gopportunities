package handle

import (
    "github.com/devsvasconcelos/gopportunities.git/internal/usecases"
)

type JobHandler struct {
	jobUseCase *usecases.JobUseCase
}

func NewJobHandler(uc *usecases.JobUseCase) *JobHandler {
	return &JobHandler{jobUseCase: uc}
}
