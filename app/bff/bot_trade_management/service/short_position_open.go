package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
	"errors"
)

func (s *botContinuingBarService) openShortPosition(ctx context.Context, request *domain.HeikinAshiDomain, tradeRequest *domain.Trade) error {
	botOpening, err := s.botOpeningService.Get(ctx, request)
	if err != nil {
		return err
	}
	if botOpening != nil {
		return errors.New("bot opening are running")
	}

	tradeRequest.SetBuySide()
	// use bn_client_id send to binance to bn_bot
	err = s.tradeService.NewOrder(ctx, tradeRequest)
	if err != nil {
		return err
	}
	return nil
}
