package routes

import (
	"github.com/Abeldlp/reservation-service/api-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeDoctorRoutes(r *gin.Engine) {
	doctor := r.Group("/doctors")

	{
		doctor.GET("", controllers.GetAllDoctors)
		doctor.GET("/:id", controllers.GetDoctor)
		doctor.POST("", controllers.CreateDoctor)
		doctor.PUT("/:id", controllers.UpdateDoctor)
		doctor.DELETE("/:id", controllers.DeleteDoctor)
	}
}
