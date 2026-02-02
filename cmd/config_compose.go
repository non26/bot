package cmd

import (
	"bot/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAWSAppLog() (*config.Config, error) {
	config, err := config.ReadAWSAppConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func UpdateConfig(g *gin.Engine, config *config.Config) {
	g.GET("/config/update", func(c *gin.Context) {
		_config, err := ReadAWSAppLog()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		config = _config
		c.JSON(http.StatusOK, gin.H{"message": "config updated"})
	})
}
