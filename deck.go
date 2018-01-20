package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

//return two values by declaring their type in parentheses
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
