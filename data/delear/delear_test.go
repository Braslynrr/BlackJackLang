package delear

import (
	"testing"

	"blackjack.com/deck"
)

// TestPeekcard calls delear.GetCard checking
// a valid card return.
func TestPeekcard(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	delear := NewDelear()
	_ = delear.GetDeck()

	_, err := delear.GetCard()
	if err != nil {
		t.Fatal("delear should be able to take a card")
	}
}

// TestGetDec calls delear.GetDeck checking a good ordered deck result
func TestGetDec(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	delear := NewDelear()
	err := delear.GetDeck()
	if err != nil {
		t.Fatalf("GetDeck() should not fail: %v", err.Error())
	}
}

// TestAddtoHand calls delear.AddtoHand checking a good process of adding
func TestAddtoHand(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	delear := NewDelear()
	_ = delear.GetDeck()
	card, _ := delear.GetCard()
	delear.AddtoHand(card)
	if len(delear.Hand) < 1 {
		t.Fatal("Delear's hand should have one card")
	}
}
