package domain

type BotDomain struct {
	BotID      string
	TemplateID string
	ClientID   string
	BnClientId string
}

func NewBotDomain() *BotDomain {
	return &BotDomain{}
}
