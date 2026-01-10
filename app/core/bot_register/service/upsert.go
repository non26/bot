package service

import (
	"bot/app/core/bot_register/domain"
	"context"
)

func (s *botTemplateService) Upsert(ctx context.Context, botTemplate *domain.BotTemplate) error {
	return s.botTemplateRepository.Upsert(ctx, botTemplate)
}
