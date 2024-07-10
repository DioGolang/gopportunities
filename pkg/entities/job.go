package entities

import "time"

type WorkingModel string

const (
	InPerson WorkingModel = "in-person"
	Hybrid   WorkingModel = "hybrid"
	Remote   WorkingModel = "remote"
)

type Job struct {
	ID                         int          `json:"id"`
	Title                      string       `json:"title"`
	Role                       string       `json:"role"`
	Description                string       `json:"description"`
	RequirementsQualifications []string     `json:"requirements_qualifications"`
	WorkingModel               WorkingModel `json:"working_model"`
	ProcessSteps               []string     `json:"process_steps"`
	Company                    string       `json:"company"`
	Location                   string       `json:"location"`
	Salary                     int          `json:"salary"`
	Link                       string       `json:"link"`
	PublishDate                time.Time    `json:"publish_date"`
	ExpirationDate             time.Time    `json:"expiration_date"`
	AdditionalInformation      string       `json:"additional_information,omitempty"`
}
