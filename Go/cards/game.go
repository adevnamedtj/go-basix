package main

import (
	"time"
)

type game struct {
	Name           string
	Type           string
	ID             string
	Started        time.Time
	DeckRemaining  Deck
	Players        map[string]user
	DeckOfHands    map[string]Deck
	DeckHandingOff map[string]Deck
	DeckUsed       Deck
}

func newGame(decCount int, gameName, gameType string) game {
	newGame := game{
		Name:          gameName,
		ID:            "egehiuhwirguiHG",
		Started:       time.Now(),
		DeckRemaining: prepareDeckForTheGame(decCount),
	}

	return newGame
}



