package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	FullName  string         `json:"full_name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"-" gorm:"column:password"`
	Bio       string         `json:"bio"`
	Photo     string         `json:"photo"`
	BirthDate time.Time      `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
