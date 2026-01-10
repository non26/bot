package service

import (
	"bot/app/core/bot_register/domain"
	"bot/app/core/bot_register/infrastructure/db"
	"context"
)

type IBotTemplateService interface {
	Get(ctx context.Context, id string) (*domain.BotTemplate, error)
	GetAll(ctx context.Context) ([]*domain.BotTemplate, error)
	Upsert(ctx context.Context, botTemplate *domain.BotTemplate) error
	Delete(ctx context.Context, id string) error
}

type botTemplateService struct {
	botTemplateRepository db.IBotTemplateRepository
}

func NewBotTemplateService(botTemplateRepository db.IBotTemplateRepository) IBotTemplateService {
	return &botTemplateService{botTemplateRepository: botTemplateRepository}
}
