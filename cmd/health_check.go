package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	healthcheckpkg "github.com/non26/tradepkg/pkg/bn/health_check"
)

func HealthCheck(app *gin.Engine, msgs string) {
	app.GET(healthcheckpkg.PATH_HEALTHCHECK, func(c *gin.Context) {
		c.JSON(http.StatusOK, healthcheckpkg.NewHealthCheckResponseWith(msgs))
	})
}
