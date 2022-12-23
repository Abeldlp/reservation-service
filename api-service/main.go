package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		user := c.GetHeader("user_id")

		c.JSON(http.StatusOK, gin.H{
			"user_id": user,
			"message": "reservation-service-pong",
		})
	})
	r.Run(":8081")
}
