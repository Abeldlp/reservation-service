package routes

import (
	"github.com/Abeldlp/reservation-service/api-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeSpecialityRoutes(r *gin.Engine) {
	speciality := r.Group("/specialities")

	{
		speciality.GET("", controllers.GetAllSpecialities)
		speciality.GET("/:id", controllers.GetSpeciality)
		speciality.POST("", controllers.CreateSpeciality)
		speciality.PUT("/:id", controllers.UpdateSpeciality)
		speciality.DELETE("/:id", controllers.DeleteSpeciality)
	}
}
