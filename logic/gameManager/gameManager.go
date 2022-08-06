package gamemanager

import (
	"sync"

	"blackjack.com/player"
	"blackjack.com/playermanager"
	"blackjack.com/room"
	"blackjack.com/roommanager"
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
