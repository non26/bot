package service

import (
	"bot/app/bff/bot_trade_management/domain"
	"context"
	"errors"
)

func (s *botContinuinCandleStickBarService) ByCandleStickCandle(ctx context.Context, request *domain.CandleStickDomain) error {
	tradeRequest := request.ToTradeDomain()
	if tradeRequest.IsLongPosition() {
		if request.OpenLongPosition() {
			openingBot, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if openingBot != nil {
				return errors.New("bot opening are running")
			}

			tradeRequest.SetBuySide()
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}

			botOpening := domain.NewBotDomain()
			botOpening.BotID = request.BotID
			botOpening.ClientID = request.ClientID
			botOpening.BnClientId = tradeRequest.BnClientID
			botOpening.Restriction = request.NumberOfBarRestriction.ToStringOfJson()
			_, err = s.botOpeningService.Update(ctx, botOpening)
			if err != nil {
				return err
			}
			return nil
		} else if request.CloseLongPosition() {
			openingBotm, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if openingBotm == nil {
				return errors.New("bot opening are not running")
			}

			tradeRequest.SetSellSide()
			tradeRequest.BnClientID = openingBotm.BnClientId
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}
			return nil
		} else if request.IsBothCandleIsGreen() {
			botOpening, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if botOpening != nil {
				candleStickRestriction := domain.NewEmptyCandleStickRestriction()
				candleStickRestriction, err := candleStickRestriction.FromStringToJson(botOpening.Restriction)
				if err != nil {
					return err
				}
				if candleStickRestriction.HasRestriction() {
					candleStickRestriction.AddCurrentBar()
					if candleStickRestriction.Continue() {
						botOpening.Restriction = candleStickRestriction.ToStringOfJson()
						_, err = s.botOpeningService.Update(ctx, botOpening)
						if err != nil {
							return err
						}
						return nil
					} else {
						tradeRequest.SetSellSide()
						tradeRequest.BnClientID = botOpening.BnClientId
						err = s.tradeService.NewOrder(ctx, tradeRequest)
						if err != nil {
							return err
						}
						return nil
					}
				}
			} else {
				return nil
			}
		}
	} else if tradeRequest.IsShortPosition() {
		if request.OpenShortPosition() {
			openingBot, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if openingBot != nil {
				return errors.New("bot opening are running")
			}

			tradeRequest.SetBuySide()
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}

			botOpening := domain.NewBotDomain()
			botOpening.BotID = request.BotID
			botOpening.ClientID = request.ClientID
			botOpening.BnClientId = tradeRequest.BnClientID
			botOpening.Restriction = request.NumberOfBarRestriction.ToStringOfJson()
			_, err = s.botOpeningService.Update(ctx, botOpening)
			if err != nil {
				return err
			}
			return nil
		} else if request.CloseShortPosition() {
			openingBot, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if openingBot == nil {
				return errors.New("bot opening are not running")
			}

			tradeRequest.SetSellSide()
			tradeRequest.BnClientID = openingBot.BnClientId
			err = s.tradeService.NewOrder(ctx, tradeRequest)
			if err != nil {
				return err
			}
			return nil
		} else if request.IsBothCandleIsRed() {

			botOpening, err := s.botOpeningService.Get(ctx, request.ToBotDomain())
			if err != nil {
				return err
			}
			if botOpening != nil {
				candleStickRestriction := domain.NewEmptyCandleStickRestriction()
				candleStickRestriction, err := candleStickRestriction.FromStringToJson(botOpening.Restriction)
				if err != nil {
					return err
				}
				candleStickRestriction.AddCurrentBar()
				if candleStickRestriction.Continue() {
					botOpening.Restriction = candleStickRestriction.ToStringOfJson()
					_, err = s.botOpeningService.Update(ctx, botOpening)
					if err != nil {
						return err
					}
				} else {
					tradeRequest.SetSellSide()
					tradeRequest.BnClientID = botOpening.BnClientId
					err = s.tradeService.NewOrder(ctx, tradeRequest)
					if err != nil {
						return err
					}
					return nil
				}
			}
			return nil
		}
	}

	return nil
}
