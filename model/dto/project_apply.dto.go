package dto

import (
	"github.com/google/uuid"
)

type ProjectApplyRequestDTO struct {
	ProjectID uuid.UUID `json:"project_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	Progress  float64   `json:"progress" validate:"required"`
}

type ProjectApplyResponseDTO struct {
	Message   string    `json:"message"`
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	UserID    uuid.UUID `json:"user_id"`
	Progress  float64   `json:"progress"`
}

type ProjectApplyGetResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Project  uuid.UUID `json:"project_id"`
	UserID   uuid.UUID `json:"user_id"`
	Progress float64   `json:"progress"`
}

type ProjectApplyGetByUserIDResponseDTO struct {
	ID       uuid.UUID             `json:"id"`
	Project  ProjectGetResponseDTO `json:"project"`
	UserID   uuid.UUID             `json:"user_id"`
	Progress float64               `json:"progress"`
}
