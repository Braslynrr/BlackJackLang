package playerapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"blackjack.com/gamemanager"
	"blackjack.com/player"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gamemanager.NewGameManager().PlayerManager.GetAll())
}

func Getplayer(c *gin.Context) {
	code := c.Param("id")
	player := gamemanager.NewGameManager().PlayerManager.FindFirtsOrDefault(func(player player.Player) bool {
		return player.Code == code
	})
	if player == nil {
		c.IndentedJSON(http.StatusNoContent, errors.New("Player doesn't exist."))
		return
	}
	c.IndentedJSON(http.StatusOK, player)
}

func NewPlayer(c *gin.Context) {
	var player *player.Player

	if err := c.BindJSON(&player); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
	}

	player, err := gamemanager.NewGameManager().PlayerManager.AddPlayer(player)

	if err == nil {

		store := ginsession.FromContext(c)
		val, _ := json.Marshal(player)
		store.Set("player", val)
		_ = store.Save()

		c.IndentedJSON(http.StatusOK, player)
		return
	}
	c.IndentedJSON(http.StatusNotAcceptable, err)
}
