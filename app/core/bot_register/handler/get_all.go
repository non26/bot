package handler

import (
	"bot/app/core/bot_register/handler/res"
	"bot/app/core/bot_register/service"
	"net/http"

	"github.com/gin-gonic/gin"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	botTemplateService service.IBotTemplateService
}

func NewGetAllHandler(botTemplateService service.IBotTemplateService) *getAllHandler {
	return &getAllHandler{botTemplateService: botTemplateService}
}

func (h *getAllHandler) Handler(c *gin.Context) {
	ctx := c.Request.Context()
	botTemplates, err := h.botTemplateService.GetAll(ctx)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		response.SendGinResponse(http.StatusInternalServerError, c)
	}
	botTemplatesResponse := []*res.GetResponse{}
	for _, botTemplate := range botTemplates {
		botTemplateResponse := res.GetResponse{}
		botTemplatesResponse = append(botTemplatesResponse, botTemplateResponse.FromDomain(botTemplate))
	}
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SubAccountNotRegisteredErrorMessage, botTemplates)
	response.SendGinResponse(http.StatusOK, c)
}
