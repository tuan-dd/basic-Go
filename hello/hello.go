package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
	list := newCard()

	// list.printAll()

	// _, remainCard := deal(list, 2)
	// takeCard, _ := slice(1, 3, list)

	// // fmt.Println(len(takeCard))
	// // fmt.Println(len(remainCard))
	// takeCard.printAll()
	// remainCard.printAll()

	// fmt.Println(list.toString())

	list.saveToFile("hello.txt")

	newList := readTofFile("hello.txt")

	newList.shuffle()
	newList.printAll()
}
