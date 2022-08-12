package dealer

import (
	"testing"

	"blackjack.com/card"
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

func TestPLayerHandValue(t *testing.T) {
	dealer := NewDealer()
	card1 := card.NewCard("diamonds", "AS", 11)
	dealer.AddtoHand(card1)
	card1 = card.NewCard("diamonds", "AS", 11)
	dealer.AddtoHand(card1)
	card1 = card.NewCard("diamonds", "AS", 11)
	dealer.AddtoHand(card1)
	card1 = card.NewCard("diamonds", "8", 8)
	dealer.AddtoHand(card1)
	card1 = card.NewCard("diamonds", "10", 10)
	dealer.AddtoHand(card1)
	if dealer.CountHandValue() != 21 {
		t.Fatalf("Hand Value should be 21, hand: %v", dealer.Hand)
	}

}
