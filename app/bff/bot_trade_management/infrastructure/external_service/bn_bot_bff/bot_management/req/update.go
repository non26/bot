package req

import "bot/app/bff/bot_trade_management/domain"

type UpdateBotOpeningRequest struct {
	BotId       string `json:"bot_id"`
	TemplateId  string `json:"template_id"`
	ClientId    string `json:"client_id"`
	BnClientId  string `json:"bn_client_id"`
	Restriction string `json:"restriction"`
}

func NewUpdateBotOpeningRequest() *UpdateBotOpeningRequest {
	return &UpdateBotOpeningRequest{}
}

func (u *UpdateBotOpeningRequest) FromDomain(domain *domain.BotDomain) {
	u.BotId = domain.BotID
	u.TemplateId = domain.TemplateID
	u.ClientId = domain.ClientID
	u.BnClientId = domain.BnClientId
	u.Restriction = domain.Restriction
}
