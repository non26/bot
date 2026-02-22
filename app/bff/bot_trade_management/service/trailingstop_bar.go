package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
	"errors"
)

var CONTINUE = "continue"
var SUCCESS = "success"

func (s *trailingStopBarService) ByTrailingStopBar(ctx context.Context, request *domain.TrailingStopBarDomain) error {
	tradeRequest := request.ToTradeDomain()
	botOpening, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
	if err != nil {
		return err
	}
	if request.IsLongPosition() {
		msg := s.ByTrailingStopBarLong(ctx, request, tradeRequest, botOpening)
		if msg != SUCCESS {
			return errors.New(msg)
		}
	} else if request.IsShortPosition() {
		msg := s.ByTrailingStopBarShort(ctx, request, tradeRequest, botOpening)
		if msg != SUCCESS {
			return errors.New(msg)
		}
	} else if request.IsBothPosition() {
		msg := s.ByTrailingStopBarLong(ctx, request, tradeRequest, botOpening)
		if msg != SUCCESS {
			if msg == CONTINUE {
				msg = s.ByTrailingStopBarShort(ctx, request, tradeRequest, botOpening)
				if msg != SUCCESS {
					return errors.New(msg)
				}
			} else {
				return errors.New(msg)
			}
		}
	}
	return nil
}
