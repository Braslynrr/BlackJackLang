package main

import (
	"fmt"

	"BlackJack.com/delear"
)

func main() {
	delear := delear.NewDelear()
	err := delear.GetDeck()
	if err != nil {
		print(err.Error())
	}
	delear.ShuffleDeck()

	fmt.Printf("Delear info is: %v\n", delear)

	cart, err := delear.GetCart()
	fmt.Printf("cart:%v\n", cart)

	cart, err = delear.GetCart()
	fmt.Printf("cart:%v\n", cart)

	fmt.Printf("There are %v carts\n", delear.Deck.CurrentCarts)

	fmt.Printf("Delear info is: %v\n", delear)
}
