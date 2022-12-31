package config

import (
	"os"

	"github.com/Abeldlp/reservation-service/api-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Doctor{}, &models.Patient{}, &models.Speciality{}, &models.Appointment{})

	DB = db

	return db
}
