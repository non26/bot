package handler

import (
	"bot/app/bff/bot_trade_management/handler/req"
	"bot/app/bff/bot_trade_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type candleStickHandler struct {
	botContinuingCandleStickBarService service.IBotContinuingCandleStickBarService
}

func NewCandleStickHandler(botContinuingCandleStickBarService service.IBotContinuingCandleStickBarService) *candleStickHandler {
	return &candleStickHandler{botContinuingCandleStickBarService: botContinuingCandleStickBarService}
}

func (h *candleStickHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.CandleStickRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}

	err = h.botContinuingCandleStickBarService.ByCandleStickCandle(ctx, req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
