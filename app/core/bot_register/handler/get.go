package handler

import (
	"bot/app/core/bot_register/handler/req"
	"bot/app/core/bot_register/handler/res"
	"bot/app/core/bot_register/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHandler struct {
	botTemplateService service.IBotTemplateService
}

func NewGetHandler(botTemplateService service.IBotTemplateService) *getHandler {
	return &getHandler{botTemplateService: botTemplateService}
}

func (h *getHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	req := req.GetRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		response.SendGinResponse(http.StatusBadRequest, c)
	}
	botTemplate, err := h.botTemplateService.Get(ctx, req.ID)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
	}
	botTemplateResponse := &res.GetResponse{}
	botTemplateResponse = botTemplateResponse.FromDomain(botTemplate)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SubAccountNotRegisteredErrorMessage, botTemplate)
	response.SendGinResponse(http.StatusOK, c)

}
