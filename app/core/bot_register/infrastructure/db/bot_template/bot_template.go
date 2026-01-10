package db

import (
	"bot/app/core/bot_register/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type botTemplateRepository struct {
	dbClient *dynamodb.Client
}

func NewBotTemplateRepository(dbClient *dynamodb.Client) db.IBotTemplateRepository {
	return &botTemplateRepository{dbClient: dbClient}
}
