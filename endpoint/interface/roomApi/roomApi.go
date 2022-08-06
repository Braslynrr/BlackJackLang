package roomapi

import (
	"net/http"

	"blackjack.com/gamemanager"
	"blackjack.com/room"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var room *room.Room
	if err := c.BindJSON(&room); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}

	room, err := gamemanager.NewGameManager().AddRoom(room)
	if err == nil {
		c.IndentedJSON(http.StatusOK, room)
		return
	}
	c.IndentedJSON(http.StatusNotAcceptable, err)
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gamemanager.NewGameManager().RoomManager.GetPublicRooms())
}
