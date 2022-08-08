package dealer

import (
	"testing"

	"blackjack.com/deck"
)

// TestPeekcard calls dealer.GetCard checking
// a valid card return.
func TestPeekcard(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	dealer := NewDealer()
	_ = dealer.GetDeck()

	_, err := dealer.GetCard()
	if err != nil {
		t.Fatal("dealer should be able to take a card")
	}
}

// TestGetDec calls dealer.GetDeck checking a good ordered deck result
func TestGetDec(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	dealer := NewDealer()
	err := dealer.GetDeck()
	if err != nil {
		t.Fatalf("GetDeck() should not fail: %v", err.Error())
	}
}

// TestAddtoHand calls dealer.AddtoHand checking a good process of adding
func TestAddtoHand(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	dealer := NewDealer()
	_ = dealer.GetDeck()
	card, _ := dealer.GetCard()
	dealer.AddtoHand(card)
	if len(dealer.Hand) < 1 {
		t.Fatal("Dealer's hand should have one card")
	}
}
