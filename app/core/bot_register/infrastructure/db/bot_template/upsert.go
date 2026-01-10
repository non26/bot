package db

import (
	"bot/app/core/bot_register/domain"
	"bot/app/core/bot_register/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (b *botTemplateRepository) Upsert(ctx context.Context, botTemplate *domain.BotTemplate) error {
	table := model.NewEmptyBotTemplate()
	table.FromDomain(botTemplate)

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetBotNameField, table.BotName)
	update_config.Set(table.GetBotTagField, botTemplate.BotTag)
	update_config.Set(table.GetDescriptionField, botTemplate.Description)
	update_config.Set(table.GetCreatedAtField, botTemplate.CreatedAt)
	update_config.Set(table.GetTemplateAttributesField, botTemplate.TemplateAttributes)
	expression := update_config.BuildExpression()
	_, err := b.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.TableName()),
		Key:                       table.GetKeyClientId(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil
}
