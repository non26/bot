package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
)

type IBotOpeningService interface {
	Get(ctx context.Context, domain *domain.HeikinAshiDomain) (*domain.BotDomain, error)
}

type botOpeningService struct {
	baseurl     string
	getEndpoint string
}

func NewBotOpeningService(baseurl string, getEndpoint string) IBotOpeningService {
	return &botOpeningService{baseurl: baseurl, getEndpoint: getEndpoint}
}
