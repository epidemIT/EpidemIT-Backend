package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectApply struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ProjectID uuid.UUID      `json:"project_id"`
	UserID    uuid.UUID      `json:"user_id"`
	Progress  float64        `json:"progress"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
