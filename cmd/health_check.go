package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(app *gin.Engine, msgs string) {
	app.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": msgs})
	})
}
