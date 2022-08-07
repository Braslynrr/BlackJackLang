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

// NewRoom creates a new room
func NewRoom(code string, password string, isprivate bool) (room *Room) {
	return &Room{Code: code, password: password, Delear: delear.NewDelear(), Players: make([]player.Player, 0, 8), isprivate: isprivate}
}

// JoinPlayer checks id a player is joinable to the room
func (room *Room) JoinPlayer(player *player.Player) (*Room, error) {
	if len(room.Players) >= 8 {
		return nil, errors.New("Imposible to Join, The Room is complete.")
	} else if room.PlayingNow {
		if room.isPlayerInTheRoom(*player) != nil {
			return room, nil
		}
		return nil, errors.New("There's a game in progess, wait for it.")
	} else if room.isPlayerInTheRoom(*player) != nil {
		return room, nil
	}
	room.Players = append(room.Players, *player)
	return room, nil
}

// NextPlayertoPlay returns the next player to play
func (room *Room) NextPlayertoPlay() *player.Player {
	for _, currentplayer := range room.Players {
		if !currentplayer.IsFinished {
			return &currentplayer
		}
	}
	return nil
}

// StartGame sets the game in a initial good state
func (room *Room) StartGame() (err error) {
	room.PlayingNow = true
	err = room.Delear.GetDeck()
	return
}

// GetDelearReady sets the delear ready to play
func (room *Room) GetDelearReady() {
	room.Delear.ShuffleDeck()
	card, _ := room.Delear.GetCard()
	room.Delear.AddtoHand(card)
	card, _ = room.Delear.GetCard()
	room.Delear.AddtoHand(card)
}

// IsPublic checks if the room is public
func (room Room) IsPublic() bool {
	return !room.isprivate
}

// CheckJoin checks if the code room is the same, with the password as well,
// unless the room is public
func (room Room) CheckJoin(r Room) bool {
	return room.Code == r.Code && (r.IsPublic() || room.password == r.password)
}

// AllowConnection finds the player joined to set it the WS connection
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

// isPlayerInTheRoom Checks if the player is in the room
func (room *Room) isPlayerInTheRoom(player player.Player) *player.Player {
	for _, ply := range room.Players {
		if ply.IsEqual(player) {
			return &ply
		}
	}
	return nil
}

func (room Room) IsEqual(rm Room) bool {
	return room.Code == rm.Code && room.password == rm.password
}
