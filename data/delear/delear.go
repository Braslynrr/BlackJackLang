package delear

import (
	"blackjack.com/card"
	"blackjack.com/deck"
)

type Delear struct {
	Deck deck.Deck   `json:"deck"`
	Hand []card.Card `json:"hand"`
}

// NewDelear creates a new empty delear
func NewDelear() (delear Delear) {
	delear = Delear{Deck: deck.Deck{}, Hand: make([]card.Card, 0)}
	return
}

// GetDeck sets to the dealer a new deck ordered
func (delear *Delear) GetDeck() error {
	deck, err := deck.NewDeck()
	if err != nil {
		return err
	}
	delear.Deck = *deck
	return err
}

// GetCard takes the first card from the delear deck
func (delear *Delear) GetCard() (card card.Card, err error) {
	card, err = delear.Deck.Peek()
	return
}

// ShuffleDeck shuffles the delear's deck
func (delear *Delear) ShuffleDeck() {
	delear.Deck.ShuffleDeck()
}

// AddtoHand adds a card to the delear's hand
func (delear *Delear) AddtoHand(card card.Card) {
	delear.Hand = append(delear.Hand, card)
}
