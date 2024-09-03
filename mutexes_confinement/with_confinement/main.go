package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func manageTicket(userChain chan int, doneChan chan bool, tickets *int) {
	for {
		select {
		case user := <-userChain:

			if *tickets > 0 {
				*tickets--
				fmt.Printf("User who got a ticket: %v\n", user)
			}
		case <-doneChan:
			fmt.Printf("Tickets remaining: %d\n", *tickets)
		}
	}
}

func buyTicket(wg *sync.WaitGroup, userChain chan int, userId int) {
	defer wg.Done()
	userChain <- userId
}

func main() {
	var tickets int = 500
	wg := &sync.WaitGroup{}
	userChain := make(chan int) // Channel for sending ticket purchase requests
	doneChan := make(chan bool) // Channel for signaling the stop
	go manageTicket(userChain, doneChan, &tickets)

	for i := 1; i <= 300; i++ {
		wg.Add(1)
		go buyTicket(wg, userChain, rand.Intn(1000000))
	}
	wg.Wait()
	doneChan <- true
	close(userChain)
}
