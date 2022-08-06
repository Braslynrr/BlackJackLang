package main

import (
	"fmt"
	"net/http"

	"blackjack.com/player"
	"blackjack.com/playerapi"
	"blackjack.com/roomapi"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/rs/cors/wrapper/gin"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		player := player.NewPlayer("01", string(msg), true)
		if err = conn.WriteJSON(player); err != nil {
			fmt.Printf("Error:%v\n", err.Error())
			return
		}
	}
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/player", playerapi.GetAll)
	router.GET("/player/:id", playerapi.Getplayer)
	router.POST("/player", playerapi.NewPlayer)
	router.GET("/room", roomapi.GetAll)
	router.POST("/room", roomapi.CreateRoom)
	router.GET("/echo", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	router.LoadHTMLFiles("websockets.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "websockets.html", nil)
	})
	router.Run("localhost:8080")
}
