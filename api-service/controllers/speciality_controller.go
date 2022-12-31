package controllers

import (
	"net/http"

	"github.com/Abeldlp/reservation-service/api-service/config"
	"github.com/Abeldlp/reservation-service/api-service/models"
	"github.com/gin-gonic/gin"
)

func GetAllSpecialities(c *gin.Context) {
	var specialities []models.Speciality

	if err := config.DB.Find(&specialities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialities)
}

func GetSpeciality(c *gin.Context) {
	var speciality models.Speciality
	specialityId := c.Param("id")

	if err := config.DB.First(&speciality, specialityId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, speciality)
}

func CreateSpeciality(c *gin.Context) {
	var speciality models.Speciality

	if err := c.ShouldBindJSON(&speciality); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&speciality).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, speciality)
}

func UpdateSpeciality(c *gin.Context) {
	var speciality models.Speciality
	specialityId := c.Param("id")

	if err := config.DB.First(&speciality, specialityId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&speciality); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&speciality).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, speciality)
}

func DeleteSpeciality(c *gin.Context) {
	var speciality models.Speciality
	specialityId := c.Param("id")

	if err := config.DB.First(&speciality, specialityId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&speciality).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Speciality deleted successfully"})
}
