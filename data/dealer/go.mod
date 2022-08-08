module blackjack.com/dealer

go 1.18

replace blackjack.com/card => ..\card

replace blackjack.com/deck => ..\deck

require (
	blackjack.com/card v0.0.0-00010101000000-000000000000
	blackjack.com/deck v0.0.0-00010101000000-000000000000
)
