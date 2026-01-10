package handler

import (
	"bot/app/core/bot_register/handler/req"
	"bot/app/core/bot_register/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type upsertHandler struct {
	botTemplateService service.IBotTemplateService
}

func NewUpsertHandler(botTemplateService service.IBotTemplateService) *upsertHandler {
	return &upsertHandler{botTemplateService: botTemplateService}
}

func (h *upsertHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.UpsertRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
	}
	err = h.botTemplateService.Upsert(ctx, req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
