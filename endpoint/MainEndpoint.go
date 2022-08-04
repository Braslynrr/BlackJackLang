package main

import (
	"fmt"

	"blackjack.com/player"
	"blackjack.com/room"
)

func main() {
	room := room.NewRoom("01", "xd")
	room.JoinPlayer(player.NewPlayer("Paco", true))
	room.StartGame()
	room.GetDelearReady()
	fmt.Printf("room is: %v", room)
}
