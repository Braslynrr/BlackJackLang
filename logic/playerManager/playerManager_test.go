package playermanager

import (
	"testing"

	"blackjack.com/player"
)

// TestAddPlayer calls playerManager.AddPlayer checking the player adding process
func TestAddPlayer(t *testing.T) {
	pManager := PlayerManager{}
	_, err := pManager.AddPlayer(player.NewPlayer("01", "name", false))
	if err != nil {
		t.Fatalf("Player should be added. Error:%v", err.Error())
	}
}

// TestAddPlayerTwice calls playerManager.AddPlayer checking a error return
func TestAddPlayerTwice(t *testing.T) {
	pManager := PlayerManager{}
	pManager.AddPlayer(player.NewPlayer("01", "name", false))
	_, err := pManager.AddPlayer(player.NewPlayer("01", "name", false))
	if err == nil {
		t.Fatalf("Player should return an error")
	}
}
