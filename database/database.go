package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // DB Instance

func InitDatabase() {
	var err error
	dsn := "postgresql://adminepidemit212121:HgjdtIbXno9CUsRHOAIDSg@epidemit-cluster-9311.8nk.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	var now time.Time
	DB.Raw("SELECT NOW()").Scan(&now)

	fmt.Println(now)
	fmt.Println("Connection Opened to Database")
}
