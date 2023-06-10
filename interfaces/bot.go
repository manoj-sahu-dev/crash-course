package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	englishBot := new(englishBot)
	spanishBot := new(spanishBot)
	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (bot englishBot) getGreeting() string {
	return fmt.Sprintf("Hello English Bot!")
}
func (bot spanishBot) getGreeting() string {
	return fmt.Sprintf("Hello Spanish Bot!")
}
