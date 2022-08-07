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

// TestAddToHandt call player.AddCardToHand checking a good process
func TestAddToHandt(t *testing.T) {
	player1 := NewPlayer("01", "name", false)
	card := card.NewCard("diamonds", "AS", 11)
	player1.AddCardToHand(*card)
	if len(player1.Hand) < 0 {
		t.Log("player should has one card.")
		t.Fail()
	}
}
