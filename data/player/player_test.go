package player

import (
	"testing"

	"blackjack.com/card"
)

// TestEqual calls player.IsRqual checking a valid bolean return
func TestEqual(t *testing.T) {
	player1, player2 := NewPlayer("01", "name", false), NewPlayer("01", "name", true)

	if !player1.IsEqual(*player2) {
		t.Log("player1 and player2 should be equal.")
		t.Fail()
	}
}

// TestAddToHandt calls player.AddCardToHand checking a good process
func TestAddToHandt(t *testing.T) {
	player1 := NewPlayer("01", "name", false)
	card := card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card)
	if len(player1.Hand) < 0 {
		t.Log("player should has one card.")
		t.Fail()
	}
}

// TestClearHand calls player.ClearHand() checking player's hand is empty
func TestClearHand(t *testing.T) {
	player1 := NewPlayer("01", "name", false)
	card1 := card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card1)
	card1 = card.NewCard("diamonds", "King", 10)
	player1.AddCardToHand(*card1)
	player1.ClearHand()
	if len(player1.Hand) != 0 {
		t.Fatalf("player's %v hand should be empty", player1)
	}
}

func TestPLayerHandValue(t *testing.T) {
	player1 := NewPlayer("01", "name", false)
	card1 := card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card1)
	card1 = card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card1)
	card1 = card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card1)
	card1 = card.NewCard("diamonds", "8", 8)
	player1.AddCardToHand(*card1)
	card1 = card.NewCard("diamonds", "10", 10)
	player1.AddCardToHand(*card1)
	if player1.CountHandValue() != 21 {
		t.Fatalf("Hand Value should be 21, hand: %v", player1.Hand)
	}

}
