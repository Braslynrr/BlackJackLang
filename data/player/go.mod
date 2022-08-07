module blackjack.com/player

go 1.18

replace blackjack.com/card => ..\card

require (
	blackjack.com/card v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.0
)
