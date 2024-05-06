package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentor struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	FullName  string         `json:"full_name"`
	Company   string         `json:"company"`
	Specialty string         `json:"specialty"`
	Email     string         `json:"email"`
	Bio       string         `json:"bio"`
	Photo     string         `json:"photo"`
	Mentees   []User         `gorm:"many2many:mentor_user;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
