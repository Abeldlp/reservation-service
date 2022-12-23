package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/Abeldlp/reservation-service/api-gateway/handlers"
	"github.com/Abeldlp/reservation-service/api-gateway/models"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

func LoginStatus(c *gin.Context) {

	requestPath := c.Request.URL.Path
	var service []string = strings.Split(requestPath, "/")

	if slices.Contains(handlers.UnauthorizedPaths, service[1]) {
		c.Next()
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing auth token"})
		c.Abort()
		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(strings.TrimSpace(token), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("user", claims)
	c.Next()
}
