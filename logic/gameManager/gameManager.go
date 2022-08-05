package gamemanager

import (
	"sync"

	"blackjack.com/playermanager"
	"blackjack.com/roommanager"
)

var lock = &sync.Mutex{}

type GameManager struct {
	rooms         roommanager.RoomManager
	playerManager playermanager.PlayerManager
}

var (
	instance *GameManager
)

func NewGameManager() *GameManager {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {

		instance = &GameManager{}
	}

	return instance
}
