package service

import (
	"bot/app/core/bot_register/domain"
	"context"
)

func (s *botTemplateService) Get(ctx context.Context, id string) (*domain.BotTemplate, error) {

	return s.botTemplateRepository.Get(ctx, id)
}
