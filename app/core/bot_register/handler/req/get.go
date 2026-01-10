package req

import "bot/app/core/bot_register/domain"

type GetRequest struct {
	ID string `json:"id"`
}

func (g *GetRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ID: g.ID,
	}
}
