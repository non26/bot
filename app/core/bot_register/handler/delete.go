package handler

import (
	"bot/app/core/bot_register/handler/req"
	"bot/app/core/bot_register/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type deleteHandler struct {
	botTemplateService service.IBotTemplateService
}

func NewDeleteHandler(botTemplateService service.IBotTemplateService) *deleteHandler {
	return &deleteHandler{botTemplateService: botTemplateService}
}

func (h *deleteHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.DeleteRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
	}
	err = h.botTemplateService.Delete(ctx, req.ID)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	response.SendGinResponse(http.StatusOK, c)
}
