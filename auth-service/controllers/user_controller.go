package controllers

import (
	"net/http"

	"github.com/Abeldlp/reservation-service/auth-service/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllUsers())
}

func Create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	models.CreateUser(user)
	c.JSON(http.StatusOK, user)
}
