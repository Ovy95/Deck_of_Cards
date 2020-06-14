package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, cards := deal(cards, 5)
	hand.print()
	fmt.Println("deck below")
	cards.print()
}
