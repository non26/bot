package main

import (
	"bot/app/bff/bot_trade_management/handler/req"
	"context"
	"encoding/json"
	"fmt"

	sqsservice "bot/cmd/aws_sqs/sqs_service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.SQSEvent) error {
	service := sqsservice.NewService()
	ctx := context.Background()
	serviceID := sqsservice.NewServiceID()
	for _, record := range event.Records {
		err := processMessage(ctx, record, service, serviceID)
		if err != nil {
			fmt.Println("error", err)
			return err
		}
		fmt.Println("done")
	}

	return nil
}

func processMessage(ctx context.Context, record events.SQSMessage, service *sqsservice.Service, serviceID *sqsservice.ServiceID) error {
	fmt.Printf("Processed message %s\n", record.Body)
	var botRequest map[string]interface{}
	err := json.Unmarshal([]byte(record.Body), &botRequest)
	if err != nil {
		return err
	}

	fmt.Println("botRequest", botRequest["service_id"])
	fmt.Println("botRequest", botRequest["request"])

	ID := botRequest["service_id"]
	stringOfPayload := botRequest["request"].(map[string]interface{})
	stringOfPayloadBytes, err := json.Marshal(stringOfPayload)
	if err != nil {
		return err
	}
	switch ID {
	case serviceID.HealthCheck: // health check
		fmt.Println("health check from sqs is complete")
		return nil
	case serviceID.BotCandle: // bot candle
		var botCandleRequest req.CandleStickRequest
		err := json.Unmarshal(stringOfPayloadBytes, &botCandleRequest)
		if err != nil {
			return err
		}
		botCandleRequestDomain := botCandleRequest.ToDomain()
		err = service.BotContinuingCandleStickBarService.ByCandleStickCandle(ctx, botCandleRequestDomain)
	case serviceID.BotHeikinAshi: // bot heikin ashii
		var botHeikinAshiRequest req.HeikinAshiRequest
		err := json.Unmarshal(stringOfPayloadBytes, &botHeikinAshiRequest)
		if err != nil {
			return err
		}
		botHeikinAshiRequestDomain := botHeikinAshiRequest.ToDomain()
		err = service.BotContinuingBarService.ByHiekinAshiCandle(ctx, botHeikinAshiRequestDomain)
	case serviceID.BotTrailingStopBar: // bot trailing stop bar
		var botTrailingStopBarRequest req.TrailingStopBarRequest
		err := json.Unmarshal(stringOfPayloadBytes, &botTrailingStopBarRequest)
		if err != nil {
			return err
		}
		botTrailingStopBarRequestDomain := botTrailingStopBarRequest.ToDomain()
		err = service.TrailingStopBarService.ByTrailingStopBar(ctx, botTrailingStopBarRequestDomain)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
