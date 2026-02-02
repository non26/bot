package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
)

type ITradeService interface {
	NewOrder(ctx context.Context, request *domain.Trade) error
}

type tradeService struct {
	baseurl          string
	newOrderEndPoint string
}

func NewTradeService(baseurl string, newOrderEndPoint string) ITradeService {
	return &tradeService{baseurl: baseurl, newOrderEndPoint: newOrderEndPoint}
}
