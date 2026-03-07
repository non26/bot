package main

import (
	externalbotmanagementservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management"
	externaltradeservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade"
	"bot/app/bff/bot_trade_management/service"
	"bot/cmd"
)

type ServiceID struct {
	BotCandle          string
	BotHeikinAshi      string
	BotTrailingStopBar string
	HealthCheck        string
}

func NewServiceID() *ServiceID {
	return &ServiceID{
		HealthCheck:        "healthcheck",
		BotCandle:          "b1",
		BotHeikinAshi:      "b2",
		BotTrailingStopBar: "b3",
	}
}

type Service struct {
	BotContinuingBarService            service.IBotContinuingBarService
	BotContinuingCandleStickBarService service.IBotContinuingCandleStickBarService
	TrailingStopBarService             service.ITrailingStopBarService
}

func NewService() *Service {
	var err error
	config, err := cmd.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

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

	botContinuingCandleStickBarService := service.NewBotContinuingCandleStickBarService(
		botTradeManagementService,
		botOpeningService,
	)

	trailingStopBarService := service.NewTrailingStopBarService(
		botTradeManagementService,
		botOpeningService,
	)

	return &Service{
		BotContinuingBarService:            botContinuingBarService,
		BotContinuingCandleStickBarService: botContinuingCandleStickBarService,
		TrailingStopBarService:             trailingStopBarService,
	}
}
