package res

import "bot/app/core/bot_register/domain"

type GetResponse struct {
	ID                 string `json:"id"`
	BotName            string `json:"bot_name"`
	BotTag             string `json:"bot_tag"`
	Description        string `json:"description"`
	CreatedAt          string `json:"created_at"`
	TemplateAttributes string `json:"template_attributes"`
}

func (g *GetResponse) FromDomain(botTemplate *domain.BotTemplate) *GetResponse {
	return &GetResponse{
		ID:                 botTemplate.ID,
		BotName:            botTemplate.BotName,
		BotTag:             botTemplate.BotTag,
		Description:        botTemplate.Description,
		CreatedAt:          botTemplate.CreatedAt,
		TemplateAttributes: botTemplate.TemplateAttributes,
	}
}
