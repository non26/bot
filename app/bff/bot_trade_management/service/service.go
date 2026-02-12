package service

import (
	"bot/app/bff/bot_trade_management/domain"
	externalbotmanagementservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management"
	externaltradeservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade"
	"context"
)

type IBotContinuingBarService interface {
	ByHiekinAshiCandle(ctx context.Context, request *domain.HeikinAshiDomain) error
}

type botContinuinHeikinAshiBarService struct {
	tradeService      externaltradeservice.ITradeService
	botOpeningService externalbotmanagementservice.IBotOpeningService
}

func NewBotContinuingHeikinAshiBarService(
	tradeService externaltradeservice.ITradeService,
	botOpeningService externalbotmanagementservice.IBotOpeningService) IBotContinuingBarService {
	return &botContinuinHeikinAshiBarService{tradeService: tradeService, botOpeningService: botOpeningService}
}

type IBotContinuingCandleStickBarService interface {
	ByCandleStickCandle(ctx context.Context, request *domain.CandleStickDomain) error
}

type botContinuinCandleStickBarService struct {
	tradeService      externaltradeservice.ITradeService
	botOpeningService externalbotmanagementservice.IBotOpeningService
}

func NewBotContinuingCandleStickBarService(
	tradeService externaltradeservice.ITradeService,
	botOpeningService externalbotmanagementservice.IBotOpeningService) IBotContinuingCandleStickBarService {
	return &botContinuinCandleStickBarService{tradeService: tradeService, botOpeningService: botOpeningService}
}
