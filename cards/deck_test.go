package main

import (
	"os"
	"testing"
)

// what to test
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//*removing any existing files
	os.Remove("_decktesting")

	//*creating a new Deck
	d := newDeck()

	//*calling the save function to save it to disk
	d.saveToFile("_decktesting")

	//*new deck from file
	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 10 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadDeck))
	}

	//*Removing the decktesting file
	os.Remove("_decktesting")
}
