package main

import "fmt"

// var deckSize int
// var some = "once"

func main() {
	/*
		printState()
		// var card string = "Ace of Spades"
		// var card = "Ace of Spades"
		card := "Ace of Spades"
		fmt.Println(card)

		var num float32 = 3.147
		fmt.Println(num)

		deckSize = 52
		fmt.Println(deckSize)

		fmt.Println(some)

		// functions
		card = newCard()
		fmt.Println(card)
	*/

	/*var cards []string = []string{"Ace of Spades", "Ace of Diamonds"}
	cards = append(cards, "Ace of hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(cards[i])
	}
	*/
	/*
		for value in cards {
			fmt.Println(value)
		}

		cards.forEach(func(value, index) {
			fmt.Println(value, index)
		})
	*/
	cards := []deck{"Ace of Spades", "Ace of Diamonds"}
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// func newCard() string {
// 	return "Ace of Diamonds"
// }
