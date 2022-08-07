package room

import (
	"fmt"
	"testing"

	"blackjack.com/deck"
	"blackjack.com/player"
)

// TestIsPublic calls room.IsPublic checking if the room is public
func TestIsPublic(t *testing.T) {
	room := NewRoom("01", "", false)
	if !room.IsPublic() {
		t.Fatal("The room should be public")
	}
	room.isprivate = true
	if room.IsPublic() {
		t.Fatal("The room should not be public")
	}
}

// TestCheckJoin calls room.CheckJoin verifying the room is joinable
func TestCheckJoin(t *testing.T) {
	room := NewRoom("01", "", false)
	otherRoom := NewRoom("01", "", false)
	if !room.CheckJoin(*otherRoom) {
		t.Fatalf("Room:%v and Room:%v shoud be joineable.", room, otherRoom)
	}

	room = NewRoom("01", "123", true)
	otherRoom = NewRoom("01", "123", true)
	if !room.CheckJoin(*otherRoom) {
		t.Fatalf("Room:%v and Room:%v shoud be joineable.", room, otherRoom)
	}

	room = NewRoom("01", "124", true)
	otherRoom = NewRoom("01", "123", true)
	if room.CheckJoin(*otherRoom) {
		t.Fatalf("Room:%v and Room:%v shoud not be joineable.", room, otherRoom)
	}
}

// TestJointPlayer calls room.JoinPlayer checking players are able to join to the room
func TestJointPlayer(t *testing.T) {
	room := NewRoom("01", "", false)
	ply := player.NewPlayer("01", "name", true)
	ServerRoom, err := room.JoinPlayer(ply)
	if err != nil {
		t.Fatalf("Player %v should join to the room :%v", ply, room)
	}

	if len(ServerRoom.Players) <= 0 {
		t.Fatalf("The Room %v should have one player", ServerRoom)
	}

	for i := 2; i < 8; i++ {
		room.JoinPlayer(player.NewPlayer(fmt.Sprint(i), "name", false))
	}
	_, err = room.JoinPlayer(ply)

	if err != nil {
		t.Fatalf("JoinPlayer should not return a error: %v, becuase is trying to reconect ", err.Error())
	}
}

// TestNotJoinPlaye calls room.JoinPlayer checking players wont be joined to the room
func TestNotJoinPlaye(t *testing.T) {
	room := NewRoom("01", "", false)
	for i := 0; i < 7; i++ {
		room.JoinPlayer(player.NewPlayer(fmt.Sprint(i), "name", false))
	}

	room.PlayingNow = true
	_, err := room.JoinPlayer(player.NewPlayer("00", "name", false))

	if err == nil {
		t.Fatal("JoinPlayer Should return an error, because is playing.")
	}
	room.PlayingNow = false
	_, err = room.JoinPlayer(player.NewPlayer("09", "name", false))

	_, err = room.JoinPlayer(player.NewPlayer("00", "name", false))

	if err == nil {
		t.Fatal("JoinPlayer Should return an error, because is full.")
	}

}

// TestStartGame calls room.StartGame verifying a good initial game state
func TestStartGame(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	room := NewRoom("01", "", false)
	ply := player.NewPlayer("01", "name", true)
	room.JoinPlayer(ply)
	err := room.StartGame()
	if err != nil {
		t.Fatalf("StartGame() shout not return an error: %v", err.Error())
	}
}
