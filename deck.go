package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of 'deck' which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//create and return a list of playing cards (an array of strings)
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cloves"}
	cardValues := []string{"ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

// deal a hand of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	sliceOfStrings := []string(d) //slice of strings
	//combine them into 1 string using Join fn
	return strings.Join(sliceOfStrings, ",")

}

//save to the file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//load a list of cards from local machine (reading from harddrive)
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		//option 1: log the err and return call to newDeck
		//option 2: log the error and quit the program
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}
