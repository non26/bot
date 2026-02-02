package route

import (
	"bot/app/bff/bot_trade_management/handler"
	externaltradeservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade"
	externalbotopeningservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bot_opening"
	"bot/app/bff/bot_trade_management/service"
	"bot/config"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, config *config.Config) {
	botTradeManagementService := externaltradeservice.NewTradeService(
		config.BNBotBFF.BotTradeManagement.BaseURL,
		config.BNBotBFF.BotTradeManagement.NewOrderEndpoint,
	)
	botOpeningService := externalbotopeningservice.NewBotOpeningService(
		config.BNBotCore.BotOpening.BaseURL,
		config.BNBotCore.BotOpening.GetEndpoint,
	)
	botContinuingBarService := service.NewBotContinuingBarService(
		botTradeManagementService,
		botOpeningService,
	)

	group := router.Group("/bot-continuing-bar")

	botContinuingBarHandler := handler.NewHeikinAshiHandler(botContinuingBarService)
	group.POST("/heikin-ashi", botContinuingBarHandler.Handler)
}
