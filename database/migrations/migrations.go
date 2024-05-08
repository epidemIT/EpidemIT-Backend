package migrations

import (
	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"fmt"

	"log"
)

func RunMigrations() {
	if database.DB == nil {
		fmt.Printf("Database connection: %v\n", database.DB)
		log.Fatal("Database connection is nil")
	}

	err := database.DB.AutoMigrate(&entity.User{}, &entity.Mentor{}, &entity.Project{}, &entity.Skill{}, &entity.ProjectApply{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Migration run successfully")
}