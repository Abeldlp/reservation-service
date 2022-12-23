package routes

import (
	"github.com/Abeldlp/reservation-service/auth-service/controllers"
	"github.com/Abeldlp/reservation-service/auth-service/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(r *gin.Engine) {
	user := r.Group("/users")

	{
		user.GET("", middlewares.UserIsLogedIn, controllers.GetAll)
	}
}
