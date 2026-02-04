package req

import "bot/app/bff/bot_trade_management/domain"

type CandleRequest struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

func (c *CandleRequest) ToDomain() *domain.CandleDomain {
	return &domain.CandleDomain{
		Open:  c.Open,
		Close: c.Close,
	}
}
