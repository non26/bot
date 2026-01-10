package req

import "bot/app/core/bot_register/domain"

type DeleteRequest struct {
	ID string `json:"id"`
}

func (d *DeleteRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ID: d.ID,
	}
}
