package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentor struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	FullName  string         `json:"full_name"`
	Company   string         `json:"company"`
	Specialty string         `json:"specialty"`
	Email     string         `json:"email" gorm:"unique"`
	Bio       string         `json:"bio"`
	Photo     string         `json:"photo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
