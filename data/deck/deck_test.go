package deck

import (
	"testing"
)

// TestPeek calls deck.Peek checking
// for a valid return value.
func TestPeek(t *testing.T) {
	DeckJson = "./deck.json"
	deck, err := NewDeck()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	_, err = deck.Peek()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if deck.CurrentCards > 51 {
		t.Log("Deck should has 51 cards")
		t.Fail()
	}
}

// TestShuffle calls deck.ShuffleDeck checking
// for a good randomized process
func TestShuffle(t *testing.T) {
	DeckJson = "./deck.json"
	deck, err := NewDeck()
	deck2, err2 := NewDeck()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if err2 != nil {
		t.Log(err2.Error())
		t.Fail()
	}
	deck.ShuffleDeck()
	var countEquals int8 = 0
	for i := 0; i < 52; i++ {
		if deck.Cards[i] == deck2.Cards[i] {
			countEquals++
		}
	}
	if countEquals > 40 {
		t.Fatal("Shuffle works wrong")
	}
}

// TestEmptyDeckPeek calls deck.ShuffleDeck checking
// error when the deck is empty
func TestEmptyDeckPeek(t *testing.T) {
	DeckJson = "./deck.json"
	deck, err := NewDeck()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	for i := 0; i < 52; i++ {
		_, err = deck.Peek()
		if err != nil {
			t.Fatalf("deck.peck should not fail : %v", err.Error())
		}
	}

	_, err = deck.Peek()
	if err == nil {
		t.Fatalf("deck.peck() shoud returns an error")
	}
}
