package domain

import "encoding/json"

type CandleStickRestriction struct {
	MaxConsecutiveBar  int `json:"max_consecutive_bar"`   // this is the max number of consecutive bar that can be used to trade
	CurrentNumberOfBar int `json:"current_number_of_bar"` // this is the current number of bar that has been used to trade
}

func NewEmptyCandleStickRestriction() *CandleStickRestriction {
	return &CandleStickRestriction{
		MaxConsecutiveBar:  0,
		CurrentNumberOfBar: 0,
	}
}

func (c *CandleStickRestriction) ToStringOfJson() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(json)
}

func (c *CandleStickRestriction) FromStringToJson(jsonString string) (*CandleStickRestriction, error) {
	candleStickRestriction := NewEmptyCandleStickRestriction()
	err := json.Unmarshal([]byte(jsonString), candleStickRestriction)
	if err != nil {
		return nil, err
	}
	return candleStickRestriction, nil
}

func (c *CandleStickRestriction) Continue() bool {
	return c.CurrentNumberOfBar < c.MaxConsecutiveBar
}

func (c *CandleStickRestriction) AddCurrentBar() {
	c.CurrentNumberOfBar++
}

func (c *CandleStickRestriction) HasRestriction() bool {
	if c.MaxConsecutiveBar <= 0 {
		return false
	}
	return true
}
