package req

import "bot/app/bff/bot_trade_management/domain"

type NewOrderRequest struct {
	TemplateId   string `json:"template_id"`
	ExchangeId   string `json:"exchange_id"`
	BotId        string `json:"bot_id"`
	TradeType    string `json:"trade_type"` // future, spot
	Symbol       string `json:"symbol"`
	Quantity     string `json:"quantity"`
	Side         string `json:"side"`          // buy, sell
	PositionSide string `json:"position_side"` // long, short
	ClientId     string `json:"client_id"`
	AccountId    string `json:"account_id"`
	BnClientID   string `json:"bn_client_id"`
}

func NewNewOrderRequest() *NewOrderRequest {
	return &NewOrderRequest{}
}

func (n *NewOrderRequest) ToDomain(d *domain.Trade) {
	n.TemplateId = d.AccountID
	n.ExchangeId = d.ExchangeID
	n.BotId = d.BotId
	n.TradeType = d.TradeType
	n.Symbol = d.Symbol
	n.Quantity = d.Quantity
	n.Side = d.Side
	n.PositionSide = d.PositionSide
	n.ClientId = d.ClientID
	n.AccountId = d.AccountID
	n.BnClientID = d.BnClientID
}
