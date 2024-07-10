package domain

import "github.com/devsvasconcelos/gopportunities.git/pkg/entities"

type JobRepository interface {
	Create(job *entities.Job) error
	GetAll() ([]*entities.Job, error)
}
