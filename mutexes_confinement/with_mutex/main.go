package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	ticker := 500
	userChain := make(chan int, ticker)
	wg := &sync.WaitGroup{}
	for userId := 1; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicker(&ticker, userChain, wg, userId)
	}
	wg.Wait()
	close(userChain)
	for userId := range userChain {
		fmt.Printf("User have ticker is %v\n", userId)
	}
}

func buyTicker(ticker *int, userChain chan int, wg *sync.WaitGroup, userId int) {
	defer wg.Done()
	mutex.Lock()
	if *ticker > 0 {
		*ticker--
		fmt.Printf("Remaining ticket is %v\n", *ticker)
		userChain <- userId
	}
	mutex.Unlock()
}
