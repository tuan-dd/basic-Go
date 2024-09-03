package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMsg(s string) {
	msg = s
	printMsg()
}

func printMsg() {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	messages := []string{"tuan", "vi", "heo"}

	for _, message := range messages {
		wg.Add(1)
		go updateMsg(message)
		wg.Wait()
	}

	fmt.Println("finish")
}
