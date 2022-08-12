package dealer

import (
	"blackjack.com/card"
	"blackjack.com/deck"
)

type Dealer struct {
	Deck deck.Deck    `json:"deck"`
	Hand []*card.Card `json:"hand"`
}

// NewDealer creates a new empty dealer
func NewDealer() (dealer Dealer) {
	dealer = Dealer{Deck: deck.Deck{}, Hand: make([]*card.Card, 0)}
	return
}

// GetDeck sets to the dealer a new deck ordered
func (dealer *Dealer) GetDeck() error {
	deck, err := deck.NewDeck()
	if err != nil {
		return err
	}
	dealer.Deck = *deck
	return err
}

// GetCard takes the first card from the dealer deck
func (dealer *Dealer) GetCard() (card *card.Card, err error) {
	card, err = dealer.Deck.Peek()
	return
}

// ShuffleDeck shuffles the dealer's deck
func (dealer *Dealer) ShuffleDeck() {
	dealer.Deck.ShuffleDeck()
}

// AddtoHand adds a card to the dealer's hand
func (dealer *Dealer) AddtoHand(card *card.Card) {
	dealer.Hand = append(dealer.Hand, card)
}

// CountHandValue counts dealer's hand value
func (delear *Dealer) CountHandValue() int8 {
	var result int8 = 0
	AS := 0
	for _, card := range delear.Hand {
		result += card.Value
		if card.ValueName == "AS" {
			AS++
		}
	}

	for AS > 0 && result > 21 {
		AS--
		result -= 10
	}
	return result
}
