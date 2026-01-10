package model

import (
	"bot/app/core/bot_register/domain"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BotTemplate struct {
	ID                 string `dynamodbav:"id" dynamodb:"id"` // primary key
	BotName            string `dynamodbav:"bot_name" dynamodb:"bot_name"`
	BotTag             string `dynamodbav:"bot_tag" dynamodb:"bot_tag"`
	Description        string `dynamodbav:"description" dynamodb:"description"`
	CreatedAt          string `dynamodbav:"created_at" dynamodb:"created_at"`
	TemplateAttributes string `dynamodbav:"template_attributes" dynamodb:"template_attributes"`
}

func NewEmptyBotTemplate() *BotTemplate {
	return &BotTemplate{}
}

func (b *BotTemplate) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ID:                 b.ID,
		BotName:            b.BotName,
		BotTag:             b.BotTag,
		Description:        b.Description,
		CreatedAt:          b.CreatedAt,
		TemplateAttributes: b.TemplateAttributes,
	}
}

func (b *BotTemplate) FromDomain(botTemplate *domain.BotTemplate) *BotTemplate {
	return &BotTemplate{
		ID:                 botTemplate.ID,
		BotName:            botTemplate.BotName,
		BotTag:             botTemplate.BotTag,
		Description:        botTemplate.Description,
		CreatedAt:          botTemplate.CreatedAt,
		TemplateAttributes: botTemplate.TemplateAttributes,
	}
}

func (b *BotTemplate) TableName() string {
	return "bot_template"
}

func (b *BotTemplate) GetKeyClientId() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: b.ID},
	}
}

func (b *BotTemplate) GetIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ID", "dynamodb")
	return v, t
}

func (b *BotTemplate) GetBotNameField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotName", "dynamodb")
	return v, t
}

func (b *BotTemplate) GetBotTagField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotTag", "dynamodb")
	return v, t
}

func (b *BotTemplate) GetDescriptionField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Description", "dynamodb")
	return v, t
}

func (b *BotTemplate) GetCreatedAtField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v, t
}

func (b *BotTemplate) GetTemplateAttributesField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "TemplateAttributes", "dynamodb")
	return v, t
}
