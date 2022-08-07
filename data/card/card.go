package card

type Card struct {
	CardType  string `json:"cardtype"`
	ValueName string `json:"valuename"`
	Value     int8   `json:"value"`
}

// NewCard creates a new card
func NewCard(cardType string, valueName string, value int8) *Card {
	return &Card{CardType: cardType, ValueName: valueName, Value: value}
}
