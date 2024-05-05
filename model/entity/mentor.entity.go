package entity

import (
	"time"

	"gorm.io/gorm"
)

type Mentor struct {
	FullName  string         `json:"full_name"`
	Company   string         `json:"company"`
	Specialty string         `json:"specialty"`
	Bio       string         `json:"bio"`
	Photo     string         `json:"photo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
