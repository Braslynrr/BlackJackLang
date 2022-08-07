package room

import (
	"errors"

	"blackjack.com/delear"
	"blackjack.com/player"
	"github.com/gorilla/websocket"
)

type Room struct {
	Code       string `json:"code"`
	password   string
	Delear     delear.Delear   `json:"delear"`
	Players    []player.Player `json:"players"`
	PlayingNow bool            `json:"playingnow"`
	isprivate  bool
}

func NewRoom(code string, password string, isprivate bool) (room Room) {
	return Room{Code: code, password: password, Delear: delear.NewDelear(), Players: make([]player.Player, 0, 8), isprivate: isprivate}
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

func (room Room) IsPublic() bool {
	return !room.isprivate
}

func (room Room) CheckJoin(r Room) bool {
	return room.Code == r.Code && (r.IsPublic() || room.password == r.password)
}

func (room *Room) AllowConnection(player player.Player, conn *websocket.Conn) bool {
	for _, serverPlayer := range room.Players {
		if serverPlayer.IsEqual(player) {
			serverPlayer.SetConeccion(conn)
			conn.WriteJSON(map[string]string{"status": "connected"})
			return true
		}
	}
	return false
}
