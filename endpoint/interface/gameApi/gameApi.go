package gameapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"blackjack.com/gamemanager"
	"blackjack.com/player"
	"blackjack.com/room"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Action string `json:"action"`
	Info   any    `json:"info"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ConnectToGame Conects the gin server with the WS server,
// also finds the current user in session
func ConnectToGame(c *gin.Context) {
	store := ginsession.FromContext(c)
	data, ok := store.Get("player")
	var player *player.Player
	jsonInfo := fmt.Sprintf("%s", data)
	value := []byte(jsonInfo)
	err := json.Unmarshal(value, &player)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}
	gameHandler(c.Writer, c.Request, *player)
}

//printError prints error to the user
func printError(connection *websocket.Conn, err error) {
	connection.WriteJSON(map[string]interface{}{"action": "notify", "status": err})
}

// connectToRoom looks if the user is alredy loged in, to allows it connecting to the room
func connectToRoom(connection *websocket.Conn, player player.Player, info any) (room.Room, bool) {
	var room room.Room
	mapstructure.Decode(info, &room)
	return room, gamemanager.NewGameManager().RoomManager.ConnectToRoom(connection, player, room)

}

// startGame sets the game in the initial status
func startGame(player player.Player, room room.Room) (ok bool, err error) {
	ok, err = gamemanager.NewGameManager().StartGame(player, room)
	return
}

// playerPlay performs the player action
func playerPlay(player player.Player, info any, room room.Room) (ok bool, err error) {
	var action string
	mapstructure.Decode(info, &action)
	return gamemanager.NewGameManager().PlayerPlay(player, room, action)
}

// gameHandler process all user request, send information to the player as well
func gameHandler(w http.ResponseWriter, r *http.Request, player player.Player) {
	conn, err := upgrader.Upgrade(w, r, nil)
	var connected bool = true
	var room room.Room
	if err == nil {
		for connected {
			// Read message from browser
			var message Message
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			jsonInfo := fmt.Sprintf("%s", msg)
			value := []byte(jsonInfo)
			json.Unmarshal(value, &message)
			message.Action = strings.ToLower(message.Action)
			switch message.Action {
			case "connect":
				room, connected = connectToRoom(conn, player, message.Info)
			case "start":
				_, err := startGame(player, room)
				if err != nil {
					printError(conn, err)
				}
			case "play":
				connected, err = playerPlay(player, message.Info, room)
				if err != nil {
					printError(conn, err)
				}
			case "kick":
			default:
				printError(conn, errors.New("Action does not exist."))
			}

		}

		conn.Close()
	}
}
