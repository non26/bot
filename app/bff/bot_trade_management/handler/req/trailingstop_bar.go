package req

import (
	"bot/app/bff/bot_trade_management/domain"
	"fmt"
)

type TrailingStopBarRequest struct {
	Candle0      *CandleRequest `json:"candle0"`
	Candle1      *CandleRequest `json:"candle1"`
	BotID        string         `json:"bot_id"`
	ClientID     string         `json:"client_id"`
	Symbol       string         `json:"symbol"`
	PositionSide string         `json:"position_side"` // LONG. SHORT, BOTh
	AmountBase   string         `json:"amount_b"`
	AccountID    string         `json:"account_id"`
	Timeframe    int            `json:"timeframe"` // send as second
}

func NewEmptyTrailingStopBarRequest() *TrailingStopBarRequest {
	return &TrailingStopBarRequest{}
}

func (t *TrailingStopBarRequest) ConvertHourToDay(hour int) string {
	days := hour / 24
	return fmt.Sprintf("%d day", days)
}

func (t *TrailingStopBarRequest) ConvertMinuteToHour(minute int) string {
	hours := minute / 60
	if hours >= 24 {
		return t.ConvertHourToDay(hours)
	}
	return fmt.Sprintf("%d hour", hours)
}

func (t *TrailingStopBarRequest) ConvertSecondToMinute(second int) string {
	minutes := second / 60
	if minutes >= 60 {
		return t.ConvertMinuteToHour(minutes)
	}
	return fmt.Sprintf("%d minute", minutes)
}

func (t *TrailingStopBarRequest) ConvertTimeFrame() string {
	return t.ConvertSecondToMinute(t.Timeframe)
}

func (t *TrailingStopBarRequest) ToDomain() *domain.TrailingStopBarDomain {
	return &domain.TrailingStopBarDomain{
		Candle0:      t.Candle0.ToDomain(),
		Candle1:      t.Candle1.ToDomain(),
		BotID:        t.BotID,
		ClientID:     t.ClientID,
		Symbol:       t.Symbol,
		PositionSide: t.PositionSide,
		AmountBase:   t.AmountBase,
		AccountID:    t.AccountID,
		Timeframe:    t.ConvertTimeFrame(),
	}
}
