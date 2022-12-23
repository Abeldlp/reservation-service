package routes

import (
	"github.com/Abeldlp/reservation-service/auth-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeAuthRoutes(r *gin.Engine) {
	auth := r.Group("")

	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
}
