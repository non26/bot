package db

import (
	"bot/app/core/bot_register/domain"
	"context"
)

type IBotTemplateRepository interface {
	Get(ctx context.Context, id string) (*domain.BotTemplate, error)
	GetAll(ctx context.Context) ([]*domain.BotTemplate, error)
	Upsert(ctx context.Context, botTemplate *domain.BotTemplate) error
	Delete(ctx context.Context, id string) error
}
