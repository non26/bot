package handler

import (
	"bot/app/bff/bot_trade_management/handler/req"
	"bot/app/bff/bot_trade_management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type TrailingStopBarHandler struct {
	trailingStopBarService service.ITrailingStopBarService
}

func NewTrailingStopBarHandler(trailingStopBarService service.ITrailingStopBarService) *TrailingStopBarHandler {
	return &TrailingStopBarHandler{trailingStopBarService: trailingStopBarService}
}

func (h *TrailingStopBarHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.TrailingStopBarRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
		return
	}
	reqDomain := req.ToDomain()
	err = h.trailingStopBarService.ByTrailingStopBar(ctx, reqDomain)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
		return
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
