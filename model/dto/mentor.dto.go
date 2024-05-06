package dto

import (
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/google/uuid"
)

type MentorGetResponseDTO struct {
	ID        uuid.UUID     `json:"id"`
	FullName  string        `json:"full_name"`
	Email     string        `json:"email"`
	Company   string        `json:"company"`
	Specialty string        `json:"specialty"`
	Bio       string        `json:"bio"`
	Photo     string        `json:"photo"`
	Mentees   []entity.User `json:"mentees"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

type MentorCreateRequestDTO struct {
	FullName  string `json:"full_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Company   string `json:"company" validate:"required"`
	Specialty string `json:"specialty" validate:"required"`
	Bio       string `json:"bio" validate:"required"`
	Photo     string `json:"photo" validate:"required"`
}

type MentorCreateResponseDTO struct {
	Message  string    `json:"message"`
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
}
