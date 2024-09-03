package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	englishGreeting := EnglishBot{}
	spanishEnglish := SpanishBot{}

	printGreeting(englishGreeting)
	printGreeting(spanishEnglish)

}

func (eb EnglishBot) getGreeting() string {
	return "Hello"
}
func (sb SpanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
