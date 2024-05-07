package dto

import (
	"time"

	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/google/uuid"
)

type ProjectRegisterRequestDTO struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Deadline    time.Time    `json:"deadline" validate:"required"`
	PartnerName string    `json:"partner_name" validate:"required"`
	PartnerDesc string    `json:"partner_description" validate:"required"`
	Skills      []uuid.UUID `json:"skills" validate:"required"`
}

type ProjectRegisterResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Message     string    `json:"message"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type ProjectGetResponseDTO struct {
	ID                 uuid.UUID      `json:"id"`
	Name               string         `json:"name"`
	ProjectDescription string         `json:"project_description"`
	Deadline           time.Time      `json:"deadline"`
	PartnerName        string         `json:"partner_name"`
	PartnerDescription string         `json:"partner_description"`
	Users              []entity.User `json:"users"`
	Skills             []entity.Skill `json:"skills"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

//assumption : this is how we apply for a project

type ProjectApplyRegisterRequestDTO struct {
	WhyChoose string `json:"why_choose" validate:"required"`
	TellUs    string `json:"tell_us" validate:"required"`
}

type ProjectApplyGetResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	Status    string    `json:"status"`
}
