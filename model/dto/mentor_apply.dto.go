package dto

import (
	"time"

	"github.com/google/uuid"
)

type MentorApplyRequestDTO struct {
	MentorID uuid.UUID `json:"mentor_id" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
}

type MentorApplyResponseDTO struct {
	Message  string    `json:"message"`
	ID       uuid.UUID `json:"id"`
	MentorID uuid.UUID `json:"mentor_id"`
	UserID   uuid.UUID `json:"user_id"`
	Date     time.Time `json:"date"`
}

type MentorApplyGetResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	MentorID uuid.UUID `json:"mentor_id"`
	UserID   uuid.UUID `json:"user_id"`
	Date     time.Time `json:"date"`
}

type MentorApplyGetByUserIDResponseDTO struct {
	ID     uuid.UUID            `json:"id"`
	Mentor MentorGetResponseDTO `json:"mentor"`
	UserID uuid.UUID            `json:"user_id"`
}
