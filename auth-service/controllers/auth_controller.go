package controllers

import (
	"net/http"
	"os"

	"github.com/Abeldlp/reservation-service/auth-service/models"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const CookieLifeTime = 36000

func Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userFromDB := models.GetUserByEmail(user.Email)
	if userFromDB.ID == 0 {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userFromDB.ID,
		"username": userFromDB.Username,
		"email":    userFromDB.Email,
	})

	jwtTokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", jwtTokenString, CookieLifeTime, "/", "*", false, true)
	c.JSON(http.StatusOK, gin.H{"token": jwtTokenString})
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(bcryptPassword)
	if err := models.CreateUser(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	jwtTokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", jwtTokenString, CookieLifeTime, "/", "*", false, true)
	c.JSON(http.StatusCreated, gin.H{"token": jwtTokenString})
}
