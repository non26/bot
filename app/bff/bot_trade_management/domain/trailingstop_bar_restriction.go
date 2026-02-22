package domain

import (
	"encoding/json"
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/shopspring/decimal"
)

type LastbarRestrictions struct {
	OpenTimeForLastBar string `json:"open_time_for_last_bar"`
	LastBarClosePrice  string `json:"last_bar_close_price"`
	LastBarOpenPrice   string `json:"last_bar_open_price"`
	TargetPositionSide string `json:"target_position_side"` // LONG, SHORT
}

func NewEmptyLastbarRestrictions() *LastbarRestrictions {
	return &LastbarRestrictions{
		OpenTimeForLastBar: "",
		LastBarClosePrice:  "",
		LastBarOpenPrice:   "",
		TargetPositionSide: "",
	}
}

func (l *LastbarRestrictions) ToStringOfJson() string {
	json, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(json)
}

func (l *LastbarRestrictions) FromStringToJson(jsonString string) (*LastbarRestrictions, error) {
	lastbarRestrictions := NewEmptyLastbarRestrictions()
	err := json.Unmarshal([]byte(jsonString), lastbarRestrictions)
	if err != nil {
		return nil, err
	}
	return lastbarRestrictions, nil
}

func (l *LastbarRestrictions) IsTargetPositionLongSide() bool {
	if strings.ToUpper(l.TargetPositionSide) == bnconstant.LONG {
		return true
	}
	return false
}

func (l *LastbarRestrictions) IsTargetPositionShortSide() bool {
	if strings.ToUpper(l.TargetPositionSide) == bnconstant.SHORT {
		return true
	}
	return false
}

func (l *LastbarRestrictions) CurrentPriceMorethanLastBarClosePrice(price string) bool {
	_price, err := decimal.NewFromString(price)
	if err != nil {
		return false
	}
	_lastBarClosePrice, err := decimal.NewFromString(l.LastBarClosePrice)
	if err != nil {
		return false
	}
	return _price.GreaterThan(_lastBarClosePrice)
}

func (l *LastbarRestrictions) CurrentPriceLessthanLastBarOpenPrice(price string) bool {
	_price, err := decimal.NewFromString(price)
	if err != nil {
		return false
	}
	_lastBarOpenPrice, err := decimal.NewFromString(l.LastBarOpenPrice)
	if err != nil {
		return false
	}
	return _price.LessThan(_lastBarOpenPrice)
}

func (l *LastbarRestrictions) CurrentPriceMorethanLastBarOpenPrice(price string) bool {
	_price, err := decimal.NewFromString(price)
	if err != nil {
		return false
	}
	_lastBarOpenPrice, err := decimal.NewFromString(l.LastBarOpenPrice)
	if err != nil {
		return false
	}
	return _price.GreaterThan(_lastBarOpenPrice)
}

func (l *LastbarRestrictions) CurrentPriceLessthanLastBarClosePrice(price string) bool {
	_price, err := decimal.NewFromString(price)
	if err != nil {
		return false
	}
	_lastBarClosePrice, err := decimal.NewFromString(l.LastBarClosePrice)
	if err != nil {
		return false
	}
	return _price.LessThan(_lastBarClosePrice)
}
