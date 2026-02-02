package handler

import (
	"bot/app/bff/bot_trade_management/handler/req"
	"bot/app/bff/bot_trade_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type heikinashiHandler struct {
	botContinuingBarService service.IBotContinuingBarService
}

func NewHeikinAshiHandler(botContinuingBarService service.IBotContinuingBarService) *heikinashiHandler {
	return &heikinashiHandler{botContinuingBarService: botContinuingBarService}
}

func (h *heikinashiHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.HeikinAshiRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
	}

	err = h.botContinuingBarService.ByHiekinAshiCandle(ctx, req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)

}
