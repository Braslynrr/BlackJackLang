module blackjack.com/room

go 1.18

replace blackjack.com/dealer => ..\dealer

replace blackjack.com/player => ..\player

require github.com/gorilla/websocket v1.5.0
