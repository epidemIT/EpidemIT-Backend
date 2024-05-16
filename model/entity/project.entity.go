package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Skill struct {
	ID          uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Project struct {
	ID                 uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Name               string         `json:"name"`
	ProjectDescription string         `json:"project_description"`
	ShortDescription   string         `json:"short_description"`
	Price              int            `json:"price"`
	MetodeBelajar      string         `json:"metode_belajar"`
	PeralatanBelajar   string         `json:"peralatan_belajar"`
	Silabus            string         `json:"silabus"`
	TotalHours         int            `json:"total_hours"`
	ImageURL           string         `json:"image_url"`
	Deadline           time.Time      `json:"deadline"`
	PartnerName        string         `json:"partner_name"`
	PartnerDescription string         `json:"partner_description"`
	Users              []User         `gorm:"many2many:user_project;"`
	Skills             []Skill        `gorm:"many2many:project_skill;"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
