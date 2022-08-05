package room

import (
	"errors"

	"blackjack.com/delear"
	"blackjack.com/player"
)

type Room struct {
	Code       string `json:"code"`
	Password   string
	Delear     delear.Delear   `json:"delear"`
	Players    []player.Player `json:"players"`
	PlayingNow bool            `json:"playingnow"`
	isprivite  bool
}

func NewRoom(code string, password string, isprivate bool) (room Room) {
	return Room{Code: code, Password: password, Delear: delear.NewDelear(), Players: make([]player.Player, 0, 8), isprivite: isprivate}
}

func (room *Room) JoinPlayer(player *player.Player) (*Room, error) {
	if len(room.Players) >= 8 {
		return nil, errors.New("Imposible to Join, The Room is complete.")
	} else if room.PlayingNow {
		return nil, errors.New("There's a game in progess, wait for it.")
	}
	room.Players = append(room.Players, *player)
	return room, nil
}

func (room *Room) NextPlayertoPlay() *player.Player {
	for _, currentplayer := range room.Players {
		if !currentplayer.IsFinished {
			return &currentplayer
		}
	}
	return nil
}

func (room *Room) StartGame() (err error) {
	room.PlayingNow = true
	err = room.Delear.GetDeck()
	return
}

func (room *Room) GetDelearReady() {
	room.Delear.ShuffleDeck()
	cart, _ := room.Delear.GetCart()
	room.Delear.AddtoHand(cart)
	cart, _ = room.Delear.GetCart()
	room.Delear.AddtoHand(cart)
}