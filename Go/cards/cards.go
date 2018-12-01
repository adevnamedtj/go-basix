package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const filePermissions os.FileMode = 0666
const dataStoreFileNameDeckOfHand string = "DataStore_DeckOfHand.json"
const dataStoreFileNameDeckOfRemaining string = "DataStore_DeckOfRemaining.json"
const dataStoreFileNameUsers string = "DataStore_Users.json"

func main() {
	// deckOfCards := newDeck()
	// deckOfHand, _ := deckOfCards.deal(4)
	// fmt.Println("Deck of Hand ", deckOfHand)
	// deckOfHand.writeDeckToFile(dataStoreFileNameDeckOfHand, filePermissions)
	// deckOfHand = Deck{}
	// deckOfHand = getDeckFromFile(dataStoreFileNameDeckOfHand)
	// fmt.Println("Deck of hand retrived ", deckOfHand)

	// deckOfHand.shuffle()
	// fmt.Println("Deck of hand shuffled ", deckOfHand)

	tej := user{
		FirstName:    "TEJ",
		LastName:     "KAL",
		UserID:       "TEJ A",
		PasswordHash: hashPassword("tempPassword123"),
		JoinedOn:     time.Now(),
	}
	sam := user{
		FirstName:    "SAM",
		LastName:     "ALMERO",
		UserID:       "SAM A",
		PasswordHash: hashPassword("tempPassword456"),
		JoinedOn:     time.Now(),
	}

	userBytes, err := json.Marshal(map[string]user{tej.UserID: tej, sam.UserID: sam})
	if err != nil {
		log.Fatal(err)
	}
	writeToFile(dataStoreFileNameUsers, userBytes, filePermissions)

	currentGame := newGame(2, "gameNight", "Sets")

	fmt.Printf("%#v", currentGame)
}
