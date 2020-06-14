package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck to have a length of 16 , but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	//set up before cleans out existing file
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	//saves the file
	loadedDeck := newDeckFromFile("_decktesting")
	// attepts to load the file created from saved function

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
