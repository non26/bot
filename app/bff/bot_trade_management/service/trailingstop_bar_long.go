package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func (s *trailingStopBarService) ByTrailingStopBarLong(ctx context.Context,
	request *domain.TrailingStopBarDomain,
	tradeRequest *domain.Trade,
	botOpening *domain.BotDomain) string {

	tradeRequest.PositionSide = bnconstant.LONG
	msg := CONTINUE
	if botOpening == nil {
		if request.OpenLongPosition() {
			tradeRequest.SetBuySide()
			tradeRequest.BnClientID = tradeRequest.CreateBnClientID(request.ClientID)
			err := s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err.Error()
			}

			request.LastBarRestrictions = domain.NewEmptyLastbarRestrictions()
			request.LastBarRestrictions.OpenTimeForLastBar = request.GetBinanceStartTimeWithinTimeframe(request.Timeframe)
			request.LastBarRestrictions.LastBarClosePrice = request.Candle0.Close
			request.LastBarRestrictions.LastBarOpenPrice = request.Candle0.Open
			request.LastBarRestrictions.TargetPositionSide = tradeRequest.PositionSide

			botOpening = domain.NewBotDomain()
			botOpening.BotID = request.BotID
			botOpening.ClientID = request.ClientID
			botOpening.BnClientId = tradeRequest.BnClientID
			botOpening.Restriction = request.LastBarRestrictions.ToStringOfJson()
			_, err = s.botOpeningService.Update(ctx, botOpening)
			if err != nil {
				return err.Error()
			}
			msg = SUCCESS
		}
	} else {
		trailingStopBarRestriction, err := request.LastBarRestrictions.FromStringToJson(botOpening.Restriction)
		if err != nil {
			return err.Error()
		}

		if !trailingStopBarRestriction.IsTargetPositionLongSide() {
			return msg
		}

		if request.Candle0.IsGreenCandle() {
			trailingStopBarRestriction.OpenTimeForLastBar = request.GetBinanceStartTimeWithinTimeframe(request.Timeframe)
			trailingStopBarRestriction.LastBarClosePrice = request.Candle0.Close
			trailingStopBarRestriction.LastBarOpenPrice = request.Candle0.Open
			trailingStopBarRestriction.TargetPositionSide = tradeRequest.PositionSide
			botOpening.Restriction = trailingStopBarRestriction.ToStringOfJson()
			_, err = s.botOpeningService.Update(ctx, botOpening)
			if err != nil {
				return err.Error()
			}
			msg = SUCCESS
		} else {
			if trailingStopBarRestriction.CurrentPriceLessthanLastBarOpenPrice(request.Candle0.Close) {
				tradeRequest.SetSellSide()
				tradeRequest.BnClientID = botOpening.BnClientId
				err = s.tradeService.NewOrder(ctx, tradeRequest)
				if err != nil {
					return err.Error()
				}
				msg = SUCCESS
			}
		}
	}
	return msg
}
