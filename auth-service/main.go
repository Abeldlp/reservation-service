package main

import (
	"github.com/Abeldlp/reservation-service/auth-service/config"
	"github.com/Abeldlp/reservation-service/auth-service/models"
	"github.com/Abeldlp/reservation-service/auth-service/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	DB = config.InitializeDatabase()

	DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.InitializeUserRoutes(r)

	r.Run(":8082")
}
