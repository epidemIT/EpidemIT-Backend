package dto

import (
	"time"

	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/google/uuid"
)

type ProjectRegisterRequestDTO struct {
	Name             string      `json:"name" validate:"required"`
	Description      string      `json:"description" validate:"required"`
	ShortDescription string      `json:"short_description" validate:"required"`
	Price            int         `json:"price" validate:"required"`
	MetodeBelajar    string      `json:"metode_belajar" validate:"required"`
	PeralatanBelajar string      `json:"peralatan_belajar" validate:"required"`
	Silabus          string      `json:"silabus" validate:"required"`
	TotalHours       int         `json:"total_hours" validate:"required"`
	Deadline         time.Time   `json:"deadline" validate:"required"`
	ImageURL         string      `json:"image_url"`
	PartnerName      string      `json:"partner_name" validate:"required"`
	PartnerDesc      string      `json:"partner_description" validate:"required"`
	Skills           []uuid.UUID `json:"skills" validate:"required"`
}

type ProjectRegisterResponseDTO struct {
	ID               uuid.UUID `json:"id"`
	Message          string    `json:"message"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Price            int       `json:"price"`
	MetodeBelajar    string    `json:"metode_belajar"`
	PeralatanBelajar string    `json:"peralatan_belajar"`
	Silabus          string    `json:"silabus"`
	TotalHours       int       `json:"total_hours"`
	Deadline         time.Time `json:"deadline"`
	ImageURL         string    `json:"image_url"`
}

type ProjectGetResponseDTO struct {
	ID                 uuid.UUID      `json:"id"`
	Name               string         `json:"name"`
	ProjectDescription string         `json:"project_description"`
	ShortDescription   string         `json:"short_description"`
	Price              int            `json:"price"`
	MetodeBelajar      string         `json:"metode_belajar"`
	PeralatanBelajar   string         `json:"peralatan_belajar"`
	Silabus            string         `json:"silabus"`
	TotalHours         int            `json:"total_hours"`
	Deadline           time.Time      `json:"deadline"`
	ImageURL           string         `json:"image_url"`
	PartnerName        string         `json:"partner_name"`
	PartnerDescription string         `json:"partner_description"`
	Users              []entity.User  `json:"users"`
	Skills             []entity.Skill `json:"skills"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}
