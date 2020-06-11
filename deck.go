package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// Underscores are used when we dont want to use
	// a variable
	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}

	}
	return cards
}

// d can be anything
func (d deck) print() {
	// d can be anything this could be like self, or this
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
