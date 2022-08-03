package main

import (
	"fmt"

	"BlackJack.com/deck"
)

func main() {
	newdeck, err := deck.NewDeck()
	if err != nil {
		fmt.Println(err)
	}
	newdeck.ShuffleDeck()
	cart, err := newdeck.Peek()
	if err == nil {
		fmt.Printf("cart is: %v", cart)
	} else {
		fmt.Println(err)
	}

}
