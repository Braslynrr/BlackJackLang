package deck

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"time"

	"BlackJack.com/cart"
)

type Deck struct {
	Carts        []cart.Cart `json:"carts"`
	currentCarts int8
}

func NewDeck() (deck *Deck, err error) {
	content, err := ioutil.ReadFile("./data/deck/deck.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &deck)
	if err != nil {
		return
	}
	deck.currentCarts = 52
	return
}

func (deck Deck) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Carts), func(i, j int) {
		deck.Carts[i], deck.Carts[j] = deck.Carts[j], deck.Carts[i]
	})
}

func (deck Deck) Peek() (cart cart.Cart, err error) {
	err = nil
	if deck.currentCarts == 0 {
		return cart, errors.New("There aren't carts")
	}
	cart = deck.Carts[0]
	deck.Carts = deck.Carts[1:len(deck.Carts)]
	deck.currentCarts--
	return
}
