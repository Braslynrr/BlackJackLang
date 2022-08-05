package playermanager

import "blackjack.com/player"

type PlayerManager struct {
	players []player.Player
}

type playerLambda func(player.Player) bool

func (manager *PlayerManager) AddPlayer(player player.Player) player.Player {
	manager.players = append(manager.players, player)
	return player
}

func (manager *PlayerManager) FindFirtsOrDefault(predicate playerLambda) *player.Player {
	for _, player := range manager.players {
		if predicate(player) {
			return &player
		}
	}
	return nil
}
