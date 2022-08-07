package gamemanager

import (
	"testing"

	"blackjack.com/player"
	"blackjack.com/room"
)

// TestSingletonInstance checks the singleton pattern of the gameManager is working correctly
func TestSingletonInstance(t *testing.T) {
	player := player.NewPlayer("01", "name", false)
	NewGameManager().AddPlayer(player)
	list := NewGameManager().PlayerManager.GetAll()
	if list[0] != player {
		t.Fatal("The players shut be alike")
	}
}

// TestAddPlayer calls gamemanager.AddPlayer checking the player should be added correctly
func TestAddPlayer(t *testing.T) {
	gm := NewGameManager()
	py := player.NewPlayer("05", "name", false)
	player, err := gm.AddPlayer(py)
	if err != nil {
		t.Fatalf("Addplayer should not return any error, error:%v", err.Error())
	}
	if py != player {
		t.Fatalf("both players should be alike")
	}
}

// TestAddRoom calls gamemanager.AddRoom checking the room should be added correctly
func TestAddRoom(t *testing.T) {
	gm := NewGameManager()
	rm := room.NewRoom("01", "", false)
	room, err := gm.AddRoom(rm)
	if err != nil {
		t.Fatalf("AddRoom should not return any error, error:%v", err.Error())
	}
	if rm != room {
		t.Fatalf("both rooms should be alike")
	}
}

// TestJoinGame calls gamemanager.JoinGame checking the player will be joined correctly
func TestJoinGame(t *testing.T) {
	gm := NewGameManager()
	py := player.NewPlayer("03", "name", false)
	gm.AddPlayer(py)
	rm := room.NewRoom("02", "", false)
	gm.AddRoom(rm)
	room, err := gm.JoinGame(py, *rm)
	if err != nil {
		t.Fatalf("JoinGame should not return any error, error:%v", err.Error())
	}
	if !room.IsEqual(*rm) {
		t.Fatalf("both rooms should be alike. %v==%v", rm, &room)
	}
	if !room.Players[0].IsEqual(*py) {
		t.Fatalf("both players should be alike")
	}
}
