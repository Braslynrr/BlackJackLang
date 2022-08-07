package deck

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"time"

	"blackjack.com/card"
)

var DeckJson string = "./data/deck/deck.json"

type Deck struct {
	Cards        []card.Card `json:"cards"`
	CurrentCards int8        `json:"currentcards"`
}

//NewDeck creates a new deck
func NewDeck() (deck *Deck, err error) {
	content, err := ioutil.ReadFile(DeckJson)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &deck)
	if err != nil {
		return
	}
	deck.CurrentCards = 52
	return
}

// ShuffleDeck randomizes the Deck
func (deck Deck) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// Peek takes one card from the deck
func (deck *Deck) Peek() (card card.Card, err error) {
	err = nil
	if deck.CurrentCards == 0 {
		return card, errors.New("There aren't cards")
	}
	card = deck.Cards[0]
	deck.Cards = deck.Cards[1:len(deck.Cards)]
	deck.CurrentCards--
	return
}
