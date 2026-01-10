package router

import (
	"bot/app/core/bot_register/handler"
	db "bot/app/core/bot_register/infrastructure/db/bot_template"
	"bot/app/core/bot_register/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, client *dynamodb.Client) {

	botTemplateRepository := db.NewBotTemplateRepository(client)
	botTemplateService := service.NewBotTemplateService(botTemplateRepository)

	getHandler := handler.NewGetHandler(botTemplateService)
	router.POST("/bot-template", getHandler.Handler)

	getAllHandler := handler.NewGetAllHandler(botTemplateService)
	router.GET("/bot-template/all", getAllHandler.Handler)

	upsertHandler := handler.NewUpsertHandler(botTemplateService)
	router.POST("/bot-template/upsert", upsertHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(botTemplateService)
	router.POST("/bot-template/delete", deleteHandler.Handler)
}
