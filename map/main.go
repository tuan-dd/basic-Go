package main

import "fmt"

func main() {

	colors := map[string]string{} // init with nil
	// test := make(map[string]string) init without nil

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, hex := range c {
		fmt.Println("this is key:", key, "with color:", hex)
	}
}
