package externalservice

import (
	"bot/app/bff/bot_trade_management/domain"
	"bot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_bff/trade/req"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *tradeService) NewOrder(ctx context.Context, _request *domain.Trade) error {
	req := req.NewNewOrderRequest()
	req.ToDomain(_request)
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", t.baseurl, t.newOrderEndPoint)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to new order: %s", response.Status)
	}

	return nil
}
