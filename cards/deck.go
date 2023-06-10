package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

type deck []string

func createDeck() deck {
	cards := deck{}
	cardFaces := []string{"Spade", "Heart", "Diamond", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, face := range cardFaces {
		for _, value := range cardValues {
			cards = append(cards, face+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println("--------------------------------")
	for i, card := range d {
		fmt.Println(i, card)

	}
}
func (d deck) Something() string {
	return "Something"
}

func (d deck) deal(total int) (deck, deck) {
	return d[0:total], d[total:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) save(filename string) error {
	byteData := []byte(d.toString())
	return ioutil.WriteFile(filename, byteData, 0766)
}
func createDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("something went wrong: " + err.Error())
		os.Exit(1)
	}
	str := strings.Split(string(bs), ", ")
	return str, err
}
