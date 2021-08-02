package main

import (
	"testing"
	"golang.org/x/tools/go/loader"
)


func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but got")
	}
	if d[0] != "ace of Spades" {
		t.Errorf("Expected ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Cloves" {
		t.Errorf("Expected Four of Cloves but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck :=newDeckFromFile(("_decktesting"))
	if len(loadedDeck) != 16{
		t.Errorf("Expected Four of Cloves but got %v", len(loadedDeck)]
	}

}
