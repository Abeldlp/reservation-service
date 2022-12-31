package controllers

import (
	"net/http"

	"github.com/Abeldlp/reservation-service/api-service/config"
	"github.com/Abeldlp/reservation-service/api-service/models"
	"github.com/gin-gonic/gin"
)

func GetAllAppointments(c *gin.Context) {
	var appointments []models.Appointment

	if err := config.DB.Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func GetAppointment(c *gin.Context) {
	var appointment models.Appointment
	var appointmentId = c.Param("id")

	if err := config.DB.First(&appointment, appointmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment
	appointmentId := c.Param("id")

	if err := config.DB.First(&appointment, appointmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func DeleteAppointment(c *gin.Context) {
	var appointment models.Appointment
	appointmentId := c.Param("id")

	if err := config.DB.First(&appointment, appointmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
