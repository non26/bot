package db

import (
	"bot/app/core/bot_register/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (b *botTemplateRepository) Delete(ctx context.Context, id string) error {
	table := model.NewEmptyBotTemplate()
	table.ID = id
	_, err := b.dbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.TableName()),
		Key:       table.GetKeyClientId(),
	})
	if err != nil {
		return err
	}
	return nil
}
