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
func printError(connection *websocket.Conn, err error) {
	connection.WriteJSON(err)
}
func connectToRoom(connection *websocket.Conn, player player.Player, info any) bool {
	var room room.Room
	mapstructure.Decode(info, &room)
	return gamemanager.NewGameManager().RoomManager.ConnectToRoom(connection, player, room)

}
func gameHandler(w http.ResponseWriter, r *http.Request, player player.Player) {
	conn, err := upgrader.Upgrade(w, r, nil)
	var connected bool = true
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
				connected = connectToRoom(conn, player, message.Info)
			case "start":
			case "play":
			case "kick":
			default:
				printError(conn, errors.New("Action does not exist."))
			}

		}

		conn.Close()
	}
}
