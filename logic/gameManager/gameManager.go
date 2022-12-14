package gamemanager

import (
	"errors"
	"strings"
	"sync"

	"blackjack.com/player"
	"blackjack.com/playermanager"
	"blackjack.com/room"
	"blackjack.com/roommanager"
	"github.com/gorilla/websocket"
)

var lock = &sync.Mutex{}

type gameManager struct {
	RoomManager   roommanager.RoomManager
	PlayerManager playermanager.PlayerManager
}

// singleton instance
var (
	instance *gameManager
)

// NewGameManager creates and returns the gameManager instance
func NewGameManager() *gameManager {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {

		instance = &gameManager{}
	}

	return instance
}

// AddPlayer adds a player to the player list
func (game *gameManager) AddPlayer(player *player.Player) (serverplayer *player.Player, err error) {
	serverplayer, err = game.PlayerManager.AddPlayer(player)
	return serverplayer, err
}

// AddRoom adds a room to the rooms list
func (game *gameManager) AddRoom(room *room.Room) (serverRoom *room.Room, err error) {
	serverRoom = game.RoomManager.AddRoom(room)
	return
}

// JoinGame joins a player into a its selected room
func (game *gameManager) JoinGame(player *player.Player, room room.Room) (room.Room, error) {
	serverRoom := game.RoomManager.FindFirtsOrDefault(roommanager.RoomPredicate(room))
	serverRoom, err := serverRoom.JoinPlayer(player)
	return *serverRoom, err
}

// ConnectToRoom connects the player with the room using the WS
func (game *gameManager) ConnectToRoom(conn *websocket.Conn, player player.Player, room room.Room) {
	game.RoomManager.ConnectToRoom(conn, player, room)
}

// StartGame runs the StarGame method of roomManager
func (game *gameManager) StartGame(player player.Player, room room.Room) (bool, error) {
	return game.RoomManager.StartGame(player, room)
}

// PlayerPlay performs the player action
func (game *gameManager) PlayerPlay(player player.Player, room room.Room, action string) (bool, error) {
	serverRoom := game.RoomManager.FindFirtsOrDefault(roommanager.RoomPredicate(room))
	if serverRoom == nil {
		return true, errors.New("unknown Room")
	}
	action = strings.ToLower(action)
	switch action {
	case "requestcard":
		return serverRoom.RequestCard(player)
	case "passturn":
		return serverRoom.PassTurn(player)
	default:
		return true, errors.New("Uknown Action")
	}
}
