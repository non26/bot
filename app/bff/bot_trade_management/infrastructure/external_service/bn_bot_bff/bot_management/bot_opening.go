package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
)

type IBotOpeningService interface {
	Get(ctx context.Context, domain *domain.BotDomain) (*domain.BotDomain, error)
	Update(ctx context.Context, domain *domain.BotDomain) (*domain.BotDomain, error)
}

type botOpeningService struct {
	baseurl        string
	getEndpoint    string
	updateEndpoint string
}

func NewBotOpeningService(baseurl string, getEndpoint string, updateEndpoint string) IBotOpeningService {
	return &botOpeningService{baseurl: baseurl, getEndpoint: getEndpoint, updateEndpoint: updateEndpoint}
}
