package game

import (
	"math/rand"
	"strconv"
)

// Suite type for enum
type Suite int

// Enum for suites
const (
	Spade   Suite = 0
	Clubs   Suite = 1
	Hearts  Suite = 2
	Diamond Suite = 3
)

// Card struct will represent a playing card
// A = 1, ..K=13
func (suite Suite) String() string {
	names := []string{"Spade", "Clubs", "Hearts", "Diamond"}
	return names[suite]
}

// Card struct
type Card struct {
	Number int
	Suite  Suite
}

func (card Card) String() string {
	var last string
	switch card.Number {
	case 1:
		last = "A"
	case 11:
		last = "K"
	case 12:
		last = "Q"
	case 13:
		last = "J"
	default:
		last = strconv.Itoa(int(card.Number))
	}

	return card.Suite.String() + last
}

// Deck store the current deck on the table
type Deck struct {
	removedCards []Card
}

// ShuffleDeck gives a new deck of 52 cards
func ShuffleDeck() Deck {
	return Deck{removedCards: make([]Card, 0)}
}

// CardRemoved check if the deck contains the card
func (deck Deck) CardRemoved(card Card) bool {
	for _, removedCard := range deck.removedCards {
		if removedCard == card {
			return true
		}
	}
	return false
}

// AddToRemovedCards removes a card from the deck
func (deck Deck) AddToRemovedCards(card Card) {
	deck.removedCards = append(deck.removedCards, card)
}

// OpenOne removes a card from the deck
func (deck Deck) OpenOne() Card {
	return deck.PickRandom()
}

// PickRandom card from the deck
func (deck Deck) PickRandom() Card {

	number := 1 + rand.Intn(13)
	suite := Suite(rand.Intn(4))
	card := Card{Number: number, Suite: suite}
	if deck.CardRemoved(card) {
		return deck.PickRandom()
	} else {
		deck.AddToRemovedCards(card)
		return card
	}
}

// DealCards returns an array containing cards for all the players
func (deck Deck) DealCards(playerCount int) [][2]Card {
	cards := make([][2]Card, playerCount)
	for i := 0; i < playerCount; i++ {
		for j := 0; j < 2; j++ {
			cards[i][j] = deck.PickRandom()
		}
	}
	return cards
}
