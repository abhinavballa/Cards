package main

import (
	"fmt"
	"strings"
)

// create a new "deck" type
type deck []string //new type of variable called deck, behaves exactly like a slice of strings
func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func newDeck() deck { //create every possible card
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	vals := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, s := range suits {
		for _, v := range vals {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	s := strings.Join([]string(d), ",")
	return s
}
