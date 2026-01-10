package db

import (
	"bot/app/core/bot_register/domain"
	"bot/app/core/bot_register/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (b *botTemplateRepository) GetAll(ctx context.Context) ([]*domain.BotTemplate, error) {
	table := model.NewEmptyBotTemplate()

	result := []*model.BotTemplate{}

	response, err := b.dbClient.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.TableName()),
	})
	if err != nil {
		return nil, err
	}
	if response.Items == nil {
		return nil, nil
	}
	if len(response.Items) == 0 {
		return nil, nil
	}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		return nil, err
	}

	items := []*domain.BotTemplate{}
	for _, item := range result {
		items = append(items, item.ToDomain())
	}
	return items, nil
}
