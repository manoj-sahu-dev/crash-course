package main

// compile using - go run main.go deck.go

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
	// cards := deck{"Ace of Spades", "Ace of Diamonds"}
	// cards.print()
	// str := cards.Something()
	// fmt.Println(str)

	// cards := createDeck()
	// cards.print()
	// hand, deal := cards.deal(5)
	// hand.print()
	// deal.print()
	// fmt.Println(deal.toString())
	// deal.save("hello.txt")

	// err := os.Chmod("hello.txt", 0777)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("all good!")
	// }

	// some, error := createDeckFromFile("hello.txt")

	// if error == nil {
	// 	some.print()
	// }
	// i := 10
	// j := 20
	// fmt.Println("before swap: ", i, " : ", j)

	// i, j = j, i
	// fmt.Println("after swap: ", i, " : ", j)
	cards := createDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Ace of Diamonds"
// }
