package main

import (
	"bot/cmd"
	"bot/config"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"bot/app/bff/bot_trade_management/route"
)

var ginLambda *ginadapter.GinLambda
var _config *config.Config

func init() {

	var err error
	_config, err = cmd.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	// dynamodb config
	// dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	// dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(_config.DynamoDB.Region, _config.DynamoDB.Endpoint)
	// dynamodbcredential := bndynamodbconfig.NewCredential(_config.DynamoDB.Ak, _config.DynamoDB.Sk)
	// var dynamodbclient *dynamodb.Client
	// if _config.IsLocal() {
	// 	dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	// } else {
	// 	dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	// }

	// gin
	app_gin := gin.Default()
	cmd.HealthCheck(app_gin, _config.HealthCheckMsg)
	// route
	route.Route(app_gin, _config)
	cmd.UpdateConfig(app_gin, _config)

	ginLambda = ginadapter.New(app_gin)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
