package dto

import (
	"github.com/google/uuid"
)

type UserRegisterRequestDTO struct {
	FullName  string `json:"full_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type UserRegisterResponseDTO struct {
	Message   string    `json:"message"`
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	BirthDate string    `json:"birth_date"`
}

type UserGetResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	BirthDate string    `json:"birth_date"`
	Bio       string    `json:"bio"`
	Photo     string    `json:"photo"`
}
