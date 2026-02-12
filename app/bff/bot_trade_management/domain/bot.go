package domain

type BotDomain struct {
	BotID       string
	TemplateID  string
	ClientID    string
	BnClientId  string
	Restriction string
}

func NewBotDomain() *BotDomain {
	return &BotDomain{}
}
