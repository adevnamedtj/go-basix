package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const cradValueString string = "Ace,Two,Three,Four,Five,Six,Seven,Eight,Nine,Ten,Jack,Queen,King"
const cardSuitString string = "Clubs,Dimonds,Hearts,Spades"

//Card a struct to repesent a card in adeck with suit and value
type Card struct {
	Suit  string
	Value string
}

func (c Card) print() {
	fmt.Println(c.Value + " of " + c.Suit)
}
func (c Card) getName() string {
	return c.Value + " of " + c.Suit
}

//Deck is array of card struct
type Deck []Card

func (d Deck) print() {
	for _, Card := range d {
		Card.print()
	}
}

func (d Deck) toJSONBytes() []byte {
	dJSON, error := json.Marshal(d)
	if error != nil {
		log.Panicln("Marshal error:", error)
	}
	return dJSON
}

func getFromJSONBytes(b []byte) Deck {
	deck := Deck{}
	error := json.Unmarshal(b, &deck)
	if error != nil {
		log.Panicln("Unmarshal error:", error)
	}
	return deck
}

func getDeckFromFile(fileName string) Deck {
	return getFromJSONBytes(readFromFile(fileName))
}

func (d Deck) writeDeckToFile(fileName string, permission os.FileMode) {
	writeToFile(fileName, d.toJSONBytes(), permission)
}

func (d Deck) shuffle() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)

	for i := range d {
		randNum := rand.Intn(len(d) - 1)

		d[i], d[randNum] = d[randNum], d[i]
	}
}

func writeToFile(fileName string, b []byte, permission os.FileMode) {

	error := ioutil.WriteFile(fileName, b, permission)
	if error != nil {
		log.Panicln("read failue to :", fileName, "error : ", error)
	}
}

func readFromFile(fileName string) []byte {

	bytes, error := ioutil.ReadFile(fileName)
	if error != nil {
		log.Panicln("read failue from :", fileName, "error : ", error)
	}
	return bytes
}

func (d Deck) deal(i int) (Deck, Deck) {
	return d[:i], d[i:]
}

// newDeck a function to initaiate a full deck of Cards as an array and return it to caller
func newDeck() Deck {

	var cards Deck
	cardSuits := strings.Split(cardSuitString, ",")
	cradValues := strings.Split(cradValueString, ",")
	for _, cardSuite := range cardSuits {
		for _, cardValue := range cradValues {
			card := Card{
				cardSuite,
				cardValue,
			}
			cards = append(cards, card)
		}
	}
	return cards
}

func prepareDeckForTheGame(numberOfDecks int) Deck {
	deckCurrent := newDeck()
	var deckMultiCurrent Deck
	for i := 1; i <= numberOfDecks; i++ {
		for _, card := range deckCurrent {
			deckMultiCurrent = append(deckMultiCurrent, card)
			fmt.Println("DECK", i, card.getName())
		}
	}
	deckMultiCurrent.shuffle()
	return deckMultiCurrent
}

func prettyprint(x interface{}) ([]byte, error) {
	b, err := json.Marshal(x)
	if err != nil {
		var out bytes.Buffer
		err := json.Indent(&out, b, "", "  ")
		if err != nil {
			return out.Bytes(), err
		}
		return nil, err
	} else {
		return nil, err
	}

}
