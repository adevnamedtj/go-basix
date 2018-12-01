package main

import (
	"os"
	"testing"
)

const deckSize int = (13 * 4)
const dataStoreFileNameTest string = "DataStore_test.json"

//TestNewDeck is a test for newDeck func of deckLoader
func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != deckSize {
		t.Errorf("Expected %v cards in new Deck, but got %v", deckSize, len(deck))
	}

	if deck[0].Value != "Ace" {
		t.Errorf("Expected Ace as a first card in the new Deck, but got %v", deck[1].Value)
	}

	if deck[len(deck)-1].Suit != "Spades" {
		t.Errorf("Expected Spades as a last card suite in the new Deck, but got %v", deck[len(deck)-1].Suit)
	}

	if deck[25].getName() != "King of Dimonds" {
		t.Errorf("Expected King of Dimonds as a twentyfifth card in the new Deck, but got %v", deck[12].getName())
	}
}

//TestWriteDeckToFileAndGetDeckFromFile is a test for deck file operation
func TestWriteDeckToFileAndGetDeckFromFile(t *testing.T) {

	os.Remove(dataStoreFileNameTest)
	deck := newDeck()
	deck.writeDeckToFile(dataStoreFileNameTest, filePermissions)
	deckRetrived := getDeckFromFile(dataStoreFileNameTest)

	if len(deck) != len(deckRetrived) {
		t.Errorf("Missmatch in length, actual deck length is %v, retrived deck length is %v", len(deck), len(deckRetrived))
	}

	if len(deck) == len(deckRetrived) {
		for i := range deck {
			if deck[i].getName() != deckRetrived[i].getName() {
				t.Errorf("Missmatch in order, actual deck card at %v is %v, but current is %v", i, deck[i].getName(), deckRetrived[i].getName())
			}
		}
	}
	os.Remove(dataStoreFileNameTest)
}
