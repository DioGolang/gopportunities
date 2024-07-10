package interfaces

import (
	"database/sql"
	"encoding/json"
	"github.com/devsvasconcelos/gopportunities.git/internal/domain"
	"github.com/devsvasconcelos/gopportunities.git/pkg/entities"
)

type JobRepository struct {
	DB *sql.DB
}

func NewJobRepository(db *sql.DB) domain.JobRepository {
	return &JobRepository{DB: db}
}

func (repo *JobRepository) Create(job *entities.Job) error {
	reqQualBytes, err := json.Marshal(job.RequirementsQualifications)
	if err != nil {
		return err
	}

	processStepsBytes, err := json.Marshal(job.ProcessSteps)
	if err != nil {
		return err
	}

	query := `INSERT INTO jobs (title, role, description, requirements_qualifications, additional_information, working_model, process_steps, company, location, salary, link, publish_date, expiration_date)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err = repo.DB.Exec(query,
		job.Title, job.Role, job.Description, reqQualBytes, job.AdditionalInformation,
		job.WorkingModel, processStepsBytes, job.Company, job.Location, job.Salary,
		job.Link, job.PublishDate, job.ExpirationDate)
	return err
}

func (repo *JobRepository) GetAll() ([]*entities.Job, error) {
	rows, err := repo.DB.Query("SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*entities.Job
	for rows.Next() {
		var job entities.Job
		var reqQualBytes []byte
		var processStepsBytes []byte

		err := rows.Scan(&job.ID, &job.Title, &job.Role, &job.Description, &reqQualBytes,
			&job.AdditionalInformation, &job.WorkingModel, &processStepsBytes, &job.Company, &job.Location,
			&job.Salary, &job.Link, &job.PublishDate, &job.ExpirationDate)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(reqQualBytes, &job.RequirementsQualifications)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(processStepsBytes, &job.ProcessSteps)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, &job)
	}
	return jobs, nil
}
