package domain

import (
	"strconv"
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"

	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
)

type TrailingStopBarDomain struct {
	Candle0             *CandleDomain
	Candle1             *CandleDomain
	BotID               string
	ClientID            string
	Symbol              string
	PositionSide        string
	AmountBase          string
	AccountID           string
	Timeframe           string
	LastBarRestrictions *LastbarRestrictions
}

func NewEmptyTrailingStopBarDomain() *TrailingStopBarDomain {
	return &TrailingStopBarDomain{}
}

func (t *TrailingStopBarDomain) ToTradeDomain() *Trade {
	return &Trade{
		AccountID:    t.AccountID,
		Symbol:       t.Symbol,
		PositionSide: t.PositionSide,
		Quantity:     t.AmountBase,
		ClientID:     t.ClientID,
		// BnClientID:   t.BnClientID,
		BotId: t.BotID,
	}
}

func (t *TrailingStopBarDomain) ToBotDomain() *BotDomain {
	return &BotDomain{
		BotID:    t.BotID,
		ClientID: t.ClientID,
	}
}

func (t *TrailingStopBarDomain) IsLongPosition() bool {
	return strings.ToUpper(t.PositionSide) == bnconstant.LONG
}

func (t *TrailingStopBarDomain) IsShortPosition() bool {
	return strings.ToUpper(t.PositionSide) == bnconstant.SHORT
}

func (t *TrailingStopBarDomain) IsBothPosition() bool {
	return strings.ToUpper(t.PositionSide) == bnconstant.BOTHPOSITIONSIDE
}

func (t *TrailingStopBarDomain) OpenLongPosition() bool {
	return t.Candle1.IsRedCandle() && t.Candle0.IsGreenCandle()
}

func (t *TrailingStopBarDomain) OpenShortPosition() bool {
	return t.Candle1.IsGreenCandle() && t.Candle0.IsRedCandle()
}

func (t *TrailingStopBarDomain) CloseLongPosition() bool {
	return t.Candle1.IsGreenCandle() && t.Candle0.IsRedCandle()
}

func (t *TrailingStopBarDomain) CloseShortPosition() bool {
	return t.Candle1.IsRedCandle() && t.Candle0.IsGreenCandle()
}

func (t *TrailingStopBarDomain) WhenBothCandleIsGreen() bool {
	return t.Candle1.IsGreenCandle() && t.Candle0.IsGreenCandle()
}

func (t *TrailingStopBarDomain) WhenBothCandleIsRed() bool {
	return t.Candle1.IsRedCandle() && t.Candle0.IsRedCandle()
}

func (t *TrailingStopBarDomain) GetBinanceStartTimeWithinTimeframe(timeframe string) string {
	unit := strings.Split(timeframe, " ")[1]
	timeframeValue := strings.Split(timeframe, " ")[0]
	switch unit {
	case "minute":
		startTime, _ := bntime.GetBinanceStartAndEndTimeInMinuteTimeFrame(timeframeValue)
		return strconv.FormatInt(bntime.GetSpecificBnTimestamp(&startTime), 10)
	case "hour":
		startTime, _ := bntime.GetBinanceStartAndEndTimeInHourTimeFrame(timeframeValue)
		return strconv.FormatInt(bntime.GetSpecificBnTimestamp(&startTime), 10)
	case "day":
		starttime, _ := bntime.GetBinanceStartAndEndTimeInDayTimeFrame(timeframeValue)
		return strconv.FormatInt(bntime.GetSpecificBnTimestamp(&starttime), 10)
	}
	return ""
}
