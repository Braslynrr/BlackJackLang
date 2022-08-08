package roomapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"blackjack.com/gamemanager"
	"blackjack.com/player"
	"blackjack.com/room"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/go-session/session"
)

func CreateRoom(c *gin.Context) {
	var room *room.Room
	if err := c.BindJSON(&room); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}

	room, err := gamemanager.NewGameManager().AddRoom(room)
	if err == nil {
		store := ginsession.FromContext(c)
		player, ok := GetSessionPlayer(store)
		if ok != nil {

			c.AbortWithStatus(404)
			return
		}
		player.IsHost = true
		room.JoinPlayer(player)

		val, _ := json.Marshal(player)
		store.Set("player", val)
		_ = store.Save()

		c.IndentedJSON(http.StatusOK, room)
		return
	}
	c.IndentedJSON(http.StatusNotAcceptable, err)
}

func JoinRoom(c *gin.Context) {
	var room *room.Room
	if err := c.BindJSON(&room); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}
	store := ginsession.FromContext(c)
	player, ok := GetSessionPlayer(store)
	if ok != nil {

		c.AbortWithStatus(404)
		return
	}
	serverRoom, err := gamemanager.NewGameManager().JoinGame(player, *room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	serverRoom.Players = nil
	c.IndentedJSON(http.StatusOK, serverRoom)
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gamemanager.NewGameManager().RoomManager.GetPublicRooms())
}

func GetSessionPlayer(store session.Store) (*player.Player, error) {
	data, _ := store.Get("player")
	var player *player.Player
	jsonInfo := fmt.Sprintf("%s", data)
	value := []byte(jsonInfo)
	ok := json.Unmarshal(value, &player)
	return player, ok
}
