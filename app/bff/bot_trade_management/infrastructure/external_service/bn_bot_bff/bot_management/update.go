package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/bot_management/req"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (b *botOpeningService) Update(ctx context.Context, domain *domain.BotDomain) (*domain.BotDomain, error) {
	req := req.NewUpdateBotOpeningRequest()
	req.FromDomain(domain)

	url := fmt.Sprintf("%s%s", b.baseurl, b.updateEndpoint)

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

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update bot opening: %s", response.Status)
	}

	return nil, nil
}
