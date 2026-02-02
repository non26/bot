package req

import "bot/app/bff/bot_trade_management/domain"

type HeikinAshiRequest struct {
	Open         string `json:"open"`
	Close        string `json:"close"`
	BotID        string `json:"bot_id"`
	ClientID     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	// Side         string `json:"side"`
	AmountBase string `json:"amount_b"`
	AccountID  string `json:"account_id"`
	// BnClinetID is auto gen from service
}

func (h *HeikinAshiRequest) ToDomain() *domain.HeikinAshiDomain {
	return &domain.HeikinAshiDomain{
		Open:         h.Open,
		Close:        h.Close,
		BotID:        h.BotID,
		ClientID:     h.ClientID,
		Symbol:       h.Symbol,
		PositionSide: h.PositionSide,
		AmountBase:   h.AmountBase,
		AccountID:    h.AccountID,
	}
}
