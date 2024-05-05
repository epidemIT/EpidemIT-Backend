package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Project struct {
	ID                 uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Name               string         `json:"name"`
	ProjectDescription string         `json:"project_description"`
	Deadline           time.Time      `json:"deadline"`
	PartnerName        string         `json:"partner_name"`
	PartnerDescription string         `json:"partner_description"`
	UserID             []User         `gorm:"many2many:user_id;"` // many to many relationship
	Skills             []string       `gorm:"many2many:project_skills;"`
	FirstMaterial      uuid.UUID      `json:"first_material"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
