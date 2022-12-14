package player

import (
	"blackjack.com/card"
	"github.com/gorilla/websocket"
)

type Player struct {
	obj        interface{}
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	IsHost     bool        `json:"ishost"`
	IsFinished bool        `json:"-"`
	Hand       []card.Card `json:"hand"`
	connection *websocket.Conn
}

// Max value of the cards of hand summed before losing
var MAXVALUE int8 = 21

// NewPlayer creates a new PLayer
func NewPlayer(code string, name string, ishost bool) (player *Player) {
	return &Player{Code: code, Name: name, IsHost: ishost, Hand: make([]card.Card, 0), IsFinished: false, connection: nil}
}

// AddCardToHand adds a card to the player's hand
func (player *Player) AddCardToHand(card card.Card) {
	player.Hand = append(player.Hand, card)
}

// SetConeccion sets a Ws connection to the player
func (player *Player) SetConeccion(connection *websocket.Conn) {
	player.connection = connection
}

// GetConeccion gets a Ws connection to the player
func (player *Player) GetConnection() *websocket.Conn {
	return player.connection
}

// IsEqual checks if two players are equal based on their Code and Name
func (player Player) IsEqual(ply Player) bool {
	return player.Code == ply.Code && player.Name == player.Name
}

// ClearHand clears player hand
func (player *Player) ClearHand() {
	player.Hand = make([]card.Card, 0)
}

// CountHandValue sums each card value and return the result
func (player Player) CountHandValue() (result int8) {
	result = 0
	AS := 0
	for _, card := range player.Hand {
		result += card.Value
		if card.ValueName == "AS" {
			AS++
		}
	}

	for AS > 0 && result > 21 {
		AS--
		result -= 10
	}

	return
}

// StillPlaying returns if the player can taking more cards
func (player Player) StillPlaying() bool {
	return player.CountHandValue() <= MAXVALUE
}
