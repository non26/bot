package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
)

func (s *botContinuinHeikinAshiBarService) ByHiekinAshiCandle(ctx context.Context, request *domain.HeikinAshiDomain) error {
	tradeRequest := request.ToTradeDomain()
	if tradeRequest.IsLongPosition() {
		if request.OpenLongPosition() { // green candle close >= open
			s.openLongPosition(ctx, request, tradeRequest)
			return nil
		} else if request.CloseLongPosition() { // red candle close < open, do sell long position
			s.closeLongPosition(ctx, request, tradeRequest)
			return nil
		}
	} else if tradeRequest.IsShortPosition() {
		if request.OpenShortPosition() { // red candle close < open, for short is to buy
			s.openShortPosition(ctx, request, tradeRequest)
			return nil
		} else if request.CloseShortPosition() { // green candle close >= open, do sell short position
			s.closeShortPosition(ctx, request, tradeRequest)
			return nil
		}
	}

	return nil
}
