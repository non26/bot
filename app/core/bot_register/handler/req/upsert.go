package req

import "bot/app/core/bot_register/domain"

type UpsertRequest struct {
	BotName            string `json:"bot_name"`
	BotTag             string `json:"bot_tag"`
	Description        string `json:"description"`
	TemplateAttributes string `json:"template_attributes"`
}

func (u *UpsertRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		BotName:            u.BotName,
		BotTag:             u.BotTag,
		Description:        u.Description,
		TemplateAttributes: u.TemplateAttributes,
	}
}
