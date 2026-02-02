package main

import (
	"bot/app/bff/bot_trade_management/route"
	"bot/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	router := gin.Default()

	route.Route(router, config)
	router.Run(fmt.Sprintf(":%d", config.Port))
}
