package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management/req"
	"bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management/res"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (b *botOpeningService) Get(ctx context.Context, domain *domain.BotDomain) (*domain.BotDomain, error) {
	req := req.NewGetBotOpeningRequest()
	req.FromDomain(domain)

	url := fmt.Sprintf("%s%s", b.baseurl, b.getEndpoint)

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(string(body))

	var responseBody res.GetResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	if responseBody.Code == appresponse.BOTNOTFOUNDCODE && response.StatusCode == http.StatusOK {
		return nil, nil
	}

	return responseBody.Data.ToDomain(), nil
	// return nil, nil
}
