package player

import "BlackJack.com/cart"

type Player struct {
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	IsHost     bool        `json:"ishost"`
	IsFinished bool        `json:"isfinished"`
	Hand       []cart.Cart `json:"hand"`
}

func newPlayer(name string, host bool) (player *Player) {
	return &Player{Name: name, IsHost: host, Hand: make([]cart.Cart, 0), IsFinished: false}
}

func (player *Player) AddCarttoHand(cart cart.Cart) {
	player.Hand = append(player.Hand, cart)
}
