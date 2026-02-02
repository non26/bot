package req

import "bot/app/bff/bot_trade_management/domain"

type GetBotOpeningRequest struct {
	BotID string `json:"bot_id"`
}

func NewGetBotOpeningRequest() *GetBotOpeningRequest {
	return &GetBotOpeningRequest{}
}

func (g *GetBotOpeningRequest) FromDomain(domain *domain.HeikinAshiDomain) {
	g.BotID = domain.BotID
}
