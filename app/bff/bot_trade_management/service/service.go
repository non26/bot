package service

import (
	"bot/app/bff/bot_trade_management/domain"
	externaltradeservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade"
	externalbotopeningservice "bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bot_opening"
	"context"
)

type IBotContinuingBarService interface {
	ByHiekinAshiCandle(ctx context.Context, request *domain.HeikinAshiDomain) error
}

type botContinuingBarService struct {
	tradeService      externaltradeservice.ITradeService
	botOpeningService externalbotopeningservice.IBotOpeningService
}

func NewBotContinuingBarService(
	tradeService externaltradeservice.ITradeService,
	botOpeningService externalbotopeningservice.IBotOpeningService) IBotContinuingBarService {
	return &botContinuingBarService{tradeService: tradeService, botOpeningService: botOpeningService}
}
