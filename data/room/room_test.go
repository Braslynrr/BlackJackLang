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
	err := room.StartGame(false)
	if err != nil {
		t.Fatalf("StartGame() should not return an error: %v", err.Error())
	}
}

// TestDealerReady calls room.GetDealerReady() checking dealer is ready
func TestDealerReady(t *testing.T) {
	room := NewRoom("01", "", false)
	room.GetDealerReady()
	if len(room.Dealer.Hand) != 2 {
		t.Fatalf("Dealer should has two cards %v", room.Dealer.Hand)
	}
	if len(room.Dealer.Deck.Cards) != int(room.Dealer.Deck.CurrentCards) {
		t.Fatalf("Both values should be alike %v==%v", len(room.Dealer.Deck.Cards), room.Dealer.Deck.CurrentCards)
	}
}

// TestRoomAreAlike calls room.IsEqual checking two rooms are alike
func TestRoomAreAlike(t *testing.T) {
	room1 := NewRoom("01", "", false)
	room2 := NewRoom("01", "", false)
	if !room1.IsEqual(*room2) {
		t.Fatalf("The rooms should be equal. %v==%v", room1, room2)
	}
}

// TestIsPlayerHost calls room.IsPlayerHost checking the player is host of the room
func TestIsPlayerHost(t *testing.T) {
	room1 := NewRoom("01", "", false)
	ply := player.NewPlayer("01", "name", true)
	room1.JoinPlayer(ply)
	if !room1.IsPlayerHost(*ply) {
		t.Fatalf("Player should be the host of the room %v", room1)
	}
}

// TestIsPlayerInTheRoom calls room.isPlayerInTheRoom checking if the player is in the room
func TestIsPlayerInTheRoom(t *testing.T) {
	room1 := NewRoom("01", "", false)
	ply := player.NewPlayer("01", "name", true)
	room1.JoinPlayer(ply)
	ply = player.NewPlayer("02", "name1", false)
	room1.JoinPlayer(ply)
	ply2 := player.NewPlayer("04", "name3", false)
	room1.JoinPlayer(ply2)
	ply = player.NewPlayer("03", "name2", false)
	room1.JoinPlayer(ply)
	if room1.isPlayerInTheRoom(*ply2) == nil {
		t.Fatalf("Player should be in the room %v", room1)
	}
}

// TestSetHands calls room.setHands checking all the players have two card in their hands
func TestSetHands(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	room1 := NewRoom("01", "", false)
	room1.Dealer.GetDeck()
	room1.GetDealerReady()
	ply := player.NewPlayer("01", "name", true)
	room1.JoinPlayer(ply)
	ply = player.NewPlayer("02", "name1", false)
	room1.JoinPlayer(ply)
	ply2 := player.NewPlayer("04", "name3", false)
	room1.JoinPlayer(ply2)
	ply = player.NewPlayer("03", "name2", false)
	room1.JoinPlayer(ply)
	err := room1.setHands(false)
	if err != nil {
		t.Fatalf("setHands(false) should not return errros. %v", err.Error())
	}

	for _, player := range room1.Players {
		if len(player.Hand) != 2 {
			t.Fatalf("Player%v should has tow cards", player)
		}
	}

}

// TestClearHands calls room.cleanHands checking it cleans all players hand
func TestClearHands(t *testing.T) {
	deck.DeckJson = "../deck/deck.json"
	room1 := NewRoom("01", "", false)
	room1.Dealer.GetDeck()
	room1.GetDealerReady()
	ply := player.NewPlayer("01", "name", true)
	room1.JoinPlayer(ply)
	ply = player.NewPlayer("02", "name1", false)
	room1.JoinPlayer(ply)
	ply2 := player.NewPlayer("04", "name3", false)
	room1.JoinPlayer(ply2)
	ply = player.NewPlayer("03", "name2", false)
	room1.JoinPlayer(ply)
	room1.setHands(false)
	room1.cleanHands()
	for _, player := range room1.Players {
		if len(player.Hand) != 0 {
			t.Fatalf("Player%v should not has cards", player)
		}
	}
}
