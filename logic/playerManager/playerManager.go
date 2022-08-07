package playermanager

import (
	"errors"
	"fmt"

	"blackjack.com/player"
)

type PlayerManager struct {
	players []*player.Player
}

// predicate for FindFirtsOrDefault
type playerLambda func(player.Player) bool

// nameFunc return a playerLambda to check the player name in FindFirtsOrDefault method
func nameFunc(name string) playerLambda {
	return func(player player.Player) bool {
		return player.Name == name
	}
}

// AddPlayer adds a player to the player list
func (manager *PlayerManager) AddPlayer(player *player.Player) (severPlayer *player.Player, err error) {
	code := len(manager.players) + 1
	if manager.FindFirtsOrDefault(nameFunc(player.Name)) != nil {
		//todo: check if the user is free

		return nil, errors.New("User already exists")
	}
	player.Code = fmt.Sprintf("0%v", code)
	severPlayer = player
	manager.players = append(manager.players, severPlayer)
	return
}

// FindFirtsOrDefault works like the functional programming 'filter' method,
// but instead of returns a list, returns the firts element or nil
// using playerLambda as predicate
func (manager *PlayerManager) FindFirtsOrDefault(predicate playerLambda) *player.Player {
	for _, player := range manager.players {
		if predicate(*player) {
			return player
		}
	}
	return nil
}

// GetAll gets the player list
func (manager *PlayerManager) GetAll() []*player.Player {
	return manager.players
}

// GetCurrentCode returns the current number of elements in the list and adds one
func (manage PlayerManager) GetCurrentCode() int {
	return len(manage.players) + 1
}
