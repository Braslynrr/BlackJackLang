package player

import (
	"blackjack.com/cart"
	"golang.org/x/net/websocket"
)

type Player struct {
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	IsHost     bool        `json:"ishost"`
	IsFinished bool        `json:"isfinished"`
	Hand       []cart.Cart `json:"hand"`
	connection *websocket.Conn
}

func NewPlayer(code string, name string, host bool) (player *Player) {
	return &Player{Code: code, Name: name, IsHost: host, Hand: make([]cart.Cart, 0), IsFinished: false, connection: nil}
}

func (player *Player) AddCarttoHand(cart cart.Cart) {
	player.Hand = append(player.Hand, cart)
}

func (player *Player) SetConeccion(connection *websocket.Conn) {
	player.connection = connection
}

func (player *Player) GetConnection() *websocket.Conn {
	return player.connection
}
