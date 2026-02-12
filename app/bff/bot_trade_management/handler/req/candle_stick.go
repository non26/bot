package req

import "bot/app/bff/bot_trade_management/domain"

type CandleStickRequest struct {
	Candle0      *CandleRequest `json:"candle0"` // a first complete candle that stay on the right most of the chart
	Candle1      *CandleRequest `json:"candle1"` // a second complete candle that stay on the right most of the chart
	BotID        string         `json:"bot_id"`
	ClientID     string         `json:"client_id"`
	Symbol       string         `json:"symbol"`
	PositionSide string         `json:"position_side"`
	// Side         string `json:"side"`
	AmountBase             string                    `json:"amount_b"`
	AccountID              string                    `json:"account_id"`
	NumberOfBarRestriction *CandleRestrictionRequest `json:"number_of_bar_restriction"`
}

func (c *CandleStickRequest) ToDomain() *domain.CandleStickDomain {
	return &domain.CandleStickDomain{
		Candle0:                c.Candle0.ToDomain(),
		Candle1:                c.Candle1.ToDomain(),
		BotID:                  c.BotID,
		ClientID:               c.ClientID,
		Symbol:                 c.Symbol,
		PositionSide:           c.PositionSide,
		AmountB:                c.AmountBase,
		AccountID:              c.AccountID,
		NumberOfBarRestriction: c.NumberOfBarRestriction.ToDomain(),
	}
}
