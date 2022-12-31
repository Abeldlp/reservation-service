package routes

import (
	"github.com/Abeldlp/reservation-service/api-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeAppointmentRoutes(r *gin.Engine) {
	appointment := r.Group("/appointments")

	{
		appointment.GET("", controllers.GetAllAppointments)
		appointment.GET("/:id", controllers.GetAppointment)
		appointment.POST("", controllers.CreateAppointment)
		appointment.PUT("/:id", controllers.UpdateAppointment)
		appointment.DELETE("/:id", controllers.DeleteAppointment)
	}
}
