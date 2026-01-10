package db

import (
	"bot/app/core/bot_register/domain"
	"bot/app/core/bot_register/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (b *botTemplateRepository) Get(ctx context.Context, id string) (*domain.BotTemplate, error) {
	table := model.NewEmptyBotTemplate()
	table.ID = id
	result := model.NewEmptyBotTemplate()
	response, err := b.dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.TableName()),
		Key:       table.GetKeyClientId(),
	})
	if err != nil {
		return nil, err
	}
	if response.Item == nil {
		return nil, nil
	}
	err = attributevalue.UnmarshalMap(response.Item, result)
	if err != nil {
		return nil, err
	}
	return result.ToDomain(), nil
}
