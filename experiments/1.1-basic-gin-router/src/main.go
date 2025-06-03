package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	router.POST("/echo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request received and confirmed",
		})
	})

	router.Run(":8080")
}