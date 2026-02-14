package req

import "bot/app/bff/bot_trade_management/domain"

type CandleRestrictionRequest struct {
	MaxConsecutiveBar  int `json:"max_consecutive_bar"`
	CurrentNumberOfBar int `json:"current_number_of_bar"`
}

func (c *CandleRestrictionRequest) ToDomain() *domain.CandleStickRestriction {
	if c == nil {
		return domain.NewEmptyCandleStickRestriction()
	}
	return &domain.CandleStickRestriction{
		MaxConsecutiveBar:  c.MaxConsecutiveBar,
		CurrentNumberOfBar: c.CurrentNumberOfBar,
	}
}
