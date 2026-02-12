package req

import "bot/app/bff/bot_trade_management/domain"

type HeikinAshiRequest struct {
	Candle0 *CandleRequest `json:"candle0"` // a first complete candle that stay on the right most of the chart
	Candle1 *CandleRequest `json:"candle1"` // a second complete candle that stay on the right most of the chart
	// Open         string         `json:"open"`
	// Close        string         `json:"close"`
	BotID        string `json:"bot_id"`
	ClientID     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	// Side         string `json:"side"`
	AmountBase string `json:"amount_b"`
	AccountID  string `json:"account_id"`
	// AmountBaseBoth *AmountBaseForBoth `json:"amount_base_both,omitempty"`
	// ClientIDBoth   *ClientIDForBoth   `json:"client_id_both,omitempty"`
}

type AmountBaseForBoth struct {
	LongAmountBase  string `json:"long_amount_base"`
	ShortAmountBase string `json:"short_amount_base"`
}

type ClientIDForBoth struct {
	LongClientID  string `json:"long_client_id"`
	ShortClientID string `json:"short_client_id"`
}

func (h *HeikinAshiRequest) ToDomain() *domain.HeikinAshiDomain {
	return &domain.HeikinAshiDomain{
		Candle0:      h.Candle0.ToDomain(),
		Candle1:      h.Candle1.ToDomain(),
		BotID:        h.BotID,
		ClientID:     h.ClientID,
		Symbol:       h.Symbol,
		PositionSide: h.PositionSide,
		AmountBase:   h.AmountBase,
		AccountID:    h.AccountID,
	}
}
