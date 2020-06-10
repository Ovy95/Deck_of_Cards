package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// d can be anything
func (d deck) print() {
	// d can be anything this could be like self, or this
	for i, card := range d {
		fmt.Println(i, card)
	}
}
