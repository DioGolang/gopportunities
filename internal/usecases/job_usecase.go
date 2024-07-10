package usecases

import (
	"github.com/devsvasconcelos/gopportunities.git/internal/domain"
	"github.com/devsvasconcelos/gopportunities.git/pkg/entities"
)

type JobUseCase struct {
	jobRepo domain.JobRepository
}

func NewJobUseCase(repo domain.JobRepository) *JobUseCase {
	return &JobUseCase{jobRepo: repo}
}

func (uc *JobUseCase) CreateJob(job *entities.Job) error {
	return uc.jobRepo.Create(job)
}

func (uc *JobUseCase) GetJobs() ([]*entities.Job, error) {
	return uc.jobRepo.GetAll()
}
