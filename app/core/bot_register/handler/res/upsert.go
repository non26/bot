package res

import "bot/app/core/bot_register/domain"

type UpsertResponse struct {
	ID string `json:"id"`
}

func (u *UpsertResponse) FromDomain(botTemplate *domain.BotTemplate) *UpsertResponse {
	return &UpsertResponse{
		ID: botTemplate.ID,
	}
}
