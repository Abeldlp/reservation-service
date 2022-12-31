package main

import (
	"github.com/Abeldlp/reservation-service/api-service/config"
	"github.com/Abeldlp/reservation-service/api-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.InitializeDatabase()

	routes.InitializeAppointmentRoutes(r)
	routes.InitializeDoctorRoutes(r)
	routes.InitializePatientRoutes(r)
	routes.InitializeSpecialityRoutes(r)

	r.Run(":8081")
}
