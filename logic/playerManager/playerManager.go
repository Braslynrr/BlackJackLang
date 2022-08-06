package playermanager

import (
	"errors"
	"fmt"

	"blackjack.com/player"
)

type PlayerManager struct {
	players []*player.Player
}

type playerLambda func(player.Player) bool

func nameFunc(name string) func(player player.Player) bool {
	return func(player player.Player) bool {
		return player.Name == name
	}
}

func (manager *PlayerManager) AddPlayer(player *player.Player) (severPlayer *player.Player, err error) {
	code := len(manager.players) + 1
	if manager.FindFirtsOrDefault(nameFunc(player.Name)) != nil {
		//todo: check if any user connection exists

		return nil, errors.New("User already exists")
	}
	player.Code = fmt.Sprintf("0%v", code)
	severPlayer = player
	manager.players = append(manager.players, severPlayer)
	return
}

func (manager *PlayerManager) FindFirtsOrDefault(predicate playerLambda) *player.Player {
	for _, player := range manager.players {
		if predicate(*player) {
			return player
		}
	}
	return nil
}

func (manager *PlayerManager) GetAll() []*player.Player {
	return manager.players
}

func (manage PlayerManager) GetCurrentCode() int {
	return len(manage.players) + 1
}
