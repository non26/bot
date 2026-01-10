package service

import (
	"bot/app/core/bot_register/domain"
	"context"
)

func (s *botTemplateService) GetAll(ctx context.Context) ([]*domain.BotTemplate, error) {
	return s.botTemplateRepository.GetAll(ctx)
}
