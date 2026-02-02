package domain

import (
	"errors"
	"strings"
	"time"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type Trade struct {
	ExchangeID   string
	TemplateID   string
	BotId        string
	TradeType    string
	Symbol       string
	Quantity     string
	Side         string
	PositionSide string
	ClientID     string
	AccountID    string
	BnClientID   string
}

func NewTrade() *Trade {
	return &Trade{}
}

func (h *Trade) IsLongPosition() bool {
	return strings.ToUpper(h.PositionSide) == bnconstant.LONG
}

func (h *Trade) IsShortPosition() bool {
	return strings.ToUpper(h.PositionSide) == bnconstant.SHORT
}

func (h *Trade) SetBuySide() error {
	if h.IsLongPosition() {
		h.Side = bnconstant.BUY
		return nil
	} else if h.IsShortPosition() { // short postion side
		h.Side = bnconstant.SELL
		return nil
	}
	return errors.New("invalid position side")
}

func (h *Trade) SetSellSide() error {
	if h.IsLongPosition() {
		h.Side = bnconstant.SELL
		return nil
	} else if h.IsShortPosition() { // short postion side
		h.Side = bnconstant.BUY
		return nil
	}
	return errors.New("invalid position side")
}

func (b *Trade) CreateBnClientID(clientID string) string {
	// clientID that send to binance can not long than 32 characters
	bnClientID := clientID + time.Now().Format("20060102_150405")
	if len(clientID) > 32 {
		clientID = clientID[:32]
	}
	return bnClientID
}

func (b *Trade) ToBotDomain() *BotDomain {
	return &BotDomain{
		BotID:      b.BotId,
		TemplateID: b.TemplateID,
		ClientID:   b.ClientID,
		BnClientId: b.BnClientID,
	}
}
