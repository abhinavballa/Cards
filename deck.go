package main

import "fmt"

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
