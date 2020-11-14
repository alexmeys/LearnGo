package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Error("Got:", len(d), "want:", 52)
	}

	if d[0] != "Ace of Spades" {
		t.Error("Got:", d[0], "Want:", "Ace of Spades")
	}

	if d[len(d) -1] != "King of Clubs" {
		t.Error("Got:", d[len(d)-1], "Want:", "King of Clubs")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Error("Got:", len(loadedDeck), "Want:", 52)
	}

	os.Remove("_decktesting")
}
