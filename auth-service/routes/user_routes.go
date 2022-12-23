package routes

import (
	"github.com/Abeldlp/reservation-service/auth-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(r *gin.Engine) {
	user := r.Group("/users")

	{
		user.GET("", controllers.GetAll)
		user.POST("", controllers.Create)
	}
}
