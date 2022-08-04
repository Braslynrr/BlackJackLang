package delear

import (
	"blackjack.com/cart"
	"blackjack.com/deck"
)

type Delear struct {
	Deck deck.Deck   `json:"deck"`
	Hand []cart.Cart `json:"hand"`
}

func NewDelear() (delear Delear) {
	delear = Delear{Deck: deck.Deck{}, Hand: make([]cart.Cart, 0)}
	return
}

func (delear *Delear) GetDeck() error {
	deck, err := deck.NewDeck()
	if err != nil {
		return err
	}
	delear.Deck = *deck
	return err
}

func (delear *Delear) GetCart() (cart cart.Cart, err error) {
	cart, err = delear.Deck.Peek()
	return
}

func (delear *Delear) ShuffleDeck() {
	delear.Deck.ShuffleDeck()
}

func (delear *Delear) AddtoHand(cart cart.Cart) {
	delear.Hand = append(delear.Hand, cart)
}
