package main

import (
	"blackjack.com/gameapi"
	"blackjack.com/playerapi"
	"blackjack.com/roomapi"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(ginsession.New())
	router.GET("/player", playerapi.GetAll)
	router.GET("/player/:id", playerapi.Getplayer)
	router.POST("/player", playerapi.NewPlayer)
	router.GET("/room", roomapi.GetAll)
	router.POST("/room", roomapi.CreateRoom)
	router.POST("/game", roomapi.JoinRoom)
	router.GET("/game", gameapi.ConnectToGame)

	router.LoadHTMLFiles("websockets.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "websockets.html", nil)
	})
	router.Run("localhost:8080")
}
