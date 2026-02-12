package res

import "bot/app/bff/bot_trade_management/domain"

type GetResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    *GetResponseData `json:"data"`
}

type GetResponseData struct {
	BotId       string `json:"bot_id"`
	TemplateId  string `json:"template_id"`
	ClientId    string `json:"client_id"`
	BnClientId  string `json:"bn_client_id"`
	Restriction string `json:"restriction"`
}

func NewGetResponse() *GetResponse {
	return &GetResponse{}
}

func (g *GetResponseData) ToDomain() *domain.BotDomain {
	if g.BotId == "" {
		return nil
	}
	return &domain.BotDomain{
		BotID:       g.BotId,
		TemplateID:  g.TemplateId,
		ClientID:    g.ClientId,
		BnClientId:  g.BnClientId,
		Restriction: g.Restriction,
	}
}
