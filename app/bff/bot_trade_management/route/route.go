package route

import (
	"bot/app/bff/bot_trade_management/handler"
	externalbotmanagementservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management"
	externaltradeservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade"
	"bot/app/bff/bot_trade_management/service"
	"bot/config"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, config *config.Config) {
	botTradeManagementService := externaltradeservice.NewTradeService(
		config.BNBotBFF.BotTradeManagement.BaseURL,
		config.BNBotBFF.BotTradeManagement.NewOrderEndpoint,
	)
	botOpeningService := externalbotmanagementservice.NewBotOpeningService(
		config.BNBotBFF.BotManagement.BaseURL,
		config.BNBotBFF.BotManagement.GetEndpoint,
		config.BNBotBFF.BotManagement.UpdateEndpoint,
	)
	botContinuingBarService := service.NewBotContinuingHeikinAshiBarService(
		botTradeManagementService,
		botOpeningService,
	)

	group := router.Group("/bot-continuing-bar")

	botContinuingBarHandler := handler.NewHeikinAshiHandler(botContinuingBarService)
	group.POST("/heikin-ashi", botContinuingBarHandler.Handler)
}
