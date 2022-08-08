package room

import (
	"errors"

	"blackjack.com/dealer"
	"blackjack.com/deck"
	"blackjack.com/player"
	"github.com/gorilla/websocket"
)

type Room struct {
	Code       string `json:"code"`
	password   string
	Dealer     dealer.Dealer    `json:"-"`
	Players    []*player.Player `json:"-"`
	PlayingNow bool             `json:"playingnow"`
	isprivate  bool
}

// NewRoom creates a new room
func NewRoom(code string, password string, isprivate bool) (room *Room) {
	return &Room{Code: code, password: password, Dealer: dealer.NewDealer(), Players: make([]*player.Player, 0, 8), isprivate: isprivate}
}

// JoinPlayer checks id a player is joinable to the room
func (room *Room) JoinPlayer(player *player.Player) (*Room, error) {
	if len(room.Players) >= 8 {
		return nil, errors.New("Imposible to Join, The Room is complete.")
	} else if room.PlayingNow {
		if room.isPlayerInTheRoom(*player) != nil {
			return room, nil
		}
		return room, errors.New("The game is in progess, please wait.")
	} else if room.isPlayerInTheRoom(*player) != nil {
		return room, nil
	}
	room.Players = append(room.Players, player)
	return room, nil
}

// NextPlayertoPlay returns the next player to play
func (room *Room) NextPlayertoPlay() *player.Player {
	for _, currentplayer := range room.Players {
		if !currentplayer.IsFinished {
			return currentplayer
		}
	}
	return nil
}

// StartGame sets the game in a initial good state
func (room *Room) StartGame(send bool) (err error) {
	room.PlayingNow = true
	err = room.Dealer.GetDeck()
	room.GetDealerReady()
	room.cleanHands()
	room.setHands(send)
	if send {

		dealer := &dealer.Dealer{
			Hand: room.Dealer.Hand,
			Deck: deck.Deck{Cards: nil, CurrentCards: room.Dealer.Deck.CurrentCards},
		}
		msg := map[string]interface{}{
			"action": "updateDealer",
			"dealer": dealer}
		room.SendMessageToAll(msg)

		room.NextPlayertoPlay().GetConnection().WriteJSON(map[string]string{"action": "play"})
	}
	return err
}

// GetDealerReady sets the dealer ready to play
func (room *Room) GetDealerReady() {
	room.Dealer.ShuffleDeck()
	card, _ := room.Dealer.GetCard()
	room.Dealer.AddtoHand(card)
	card, _ = room.Dealer.GetCard()
	room.Dealer.AddtoHand(card)

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
			return ply
		}
	}
	return nil
}

// IsEqual checks if two rooms are alike
func (room Room) IsEqual(rm Room) bool {
	return room.Code == rm.Code && (room.IsPublic() || room.password == rm.password)
}

// cleanHands cleans all player's hands
func (room *Room) cleanHands() {

	for _, player := range room.Players {
		player.ClearHand()
	}

}

// IsPlayerHost verifies if player is host of the room
func (room *Room) IsPlayerHost(player player.Player) bool {
	for _, ply := range room.Players {
		if ply.IsEqual(player) {
			return ply.IsHost
		}
	}
	return false
}

// SendMessageToAll sends a message to all players
func (room Room) SendMessageToAll(msg map[string]interface{}) {
	for _, player := range room.Players {
		player.GetConnection().WriteJSON(msg)
	}
}

// give two carts to each player
func (room *Room) setHands(send bool) error {
	for _, player := range room.Players {
		cart, err := room.Dealer.GetCard()
		if err != nil {
			return err
		}
		player.AddCardToHand(*cart)
		cart, err = room.Dealer.GetCard()
		if err != nil {
			return err
		}
		player.AddCardToHand(*cart)
		if send {
			msg := map[string]interface{}{"action": "updateHand", "hand": player.Hand}
			player.GetConnection().WriteJSON(msg)
		}
	}
	return nil
}
