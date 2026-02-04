package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
	"errors"
)

func (s *botContinuingBarService) ByHiekinAshiCandle(ctx context.Context, request *domain.HeikinAshiDomain) error {
	tradeRequest := request.ToTradeDomain()
	if tradeRequest.IsLongPosition() {
		if request.OpenLongPosition() { // green candle close >= open
			botOpening, err := s.botOpeningService.Get(ctx, request)
			if err != nil {
				return err
			}
			if botOpening != nil {
				return errors.New("bot opening are running")
			}

			/// case when nil, means bot opening are not running			botOpening = domain.NewBotDomain()
			tradeRequest.SetBuySide()
			// use bn_client_id send to binance to bn_bot
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}
		} else if request.CloseLongPosition() { // red candle close < open, do sell long position
			botOpening, err := s.botOpeningService.Get(ctx, request)
			if err != nil {
				return err
			}
			if botOpening == nil {
				return errors.New("bot opening are not running")
			}

			tradeRequest.BnClientID = botOpening.BnClientId
			tradeRequest.SetSellSide()
			// use bn_client_id send to binance to bn_bot
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}
		}
	} else {
		if request.OpenShortPosition() { // red candle close < open, for short is to buy
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
		} else if request.CloseShortPosition() { // green candle close >= open, do sell short position
			botOpening, err := s.botOpeningService.Get(ctx, request)
			if err != nil {
				return err
			}
			if botOpening == nil {
				return errors.New("bot opening are not running")
			}
			tradeRequest.BnClientID = botOpening.BnClientId
			tradeRequest.SetSellSide()
			// use bn_client_id send to binance to bn_bot
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
