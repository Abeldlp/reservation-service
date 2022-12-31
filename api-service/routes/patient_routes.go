package routes

import (
	"github.com/Abeldlp/reservation-service/api-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializePatientRoutes(r *gin.Engine) {
	patient := r.Group("/patients")

	{
		patient.GET("", controllers.GetAllPatients)
		patient.GET("/:id", controllers.GetPatient)
		patient.POST("", controllers.CreatePatient)
		patient.PUT("/:id", controllers.UpdatePatient)
		patient.DELETE("/:id", controllers.DeletePatient)
	}
}
