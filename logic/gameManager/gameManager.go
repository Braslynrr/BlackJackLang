package gamemanager

import (
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

var (
	instance *gameManager
)

func NewGameManager() *gameManager {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {

		instance = &gameManager{}
	}

	return instance
}

func (game *gameManager) AddPlayer(player *player.Player) error {
	game.PlayerManager.AddPlayer(player)
	return nil
}

func (game *gameManager) AddRoom(room *room.Room) (serverRoom *room.Room, err error) {
	serverRoom = game.RoomManager.AddRoom(room)
	return
}

func (game *gameManager) JoinGame(player *player.Player, room room.Room) (room.Room, error) {
	serverRoom := game.RoomManager.FindFirtsOrDefault(roommanager.RoomPredicate(room))
	serverRoom, err := serverRoom.JoinPlayer(player)
	return *serverRoom, err
}

func (game *gameManager) ConnectToRoom(conn *websocket.Conn, player player.Player, room room.Room) {

	game.RoomManager.ConnectToRoom(conn, player, room)
}
