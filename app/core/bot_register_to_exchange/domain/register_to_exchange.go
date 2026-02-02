package domain

type RegisterToExchange struct {
	ExchangeId    string
	BotTemplateId string
	BotName       string
	Description   string
	CreatedAt     string
	AllowFutures  bool
	AllowSpot     bool
}
