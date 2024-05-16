package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MentorApply struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	MentorID  uuid.UUID      `json:"mentor_id"`
	UserID    uuid.UUID      `json:"user_id"`
	Date      time.Time      `json:"date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
