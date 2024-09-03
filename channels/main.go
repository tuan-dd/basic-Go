package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	links := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://www.google.com/search?q=golang",
		// "https://www.google.com/search?q=golang",
		// "https://www.google.com/search?q=golang",
		// "https://www.google.com/search?q=golang",
	}

	c := make(chan string)

	for idx, link := range links {
		// wg.Add(1)
		// go func(link string) {
		// 	// defer wg.Done()
		// 	checkStatus(link, idx, c)
		// }(link)

		go checkStatus(link, idx, c)
	}
	// wg.Wait()

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkStatus(l, 0, c)
		}()
	}

	// fmt.Println(<-c)
}

func checkStatus(link string, idx int, c chan string) {
	time.Sleep(2 * time.Second)
	res, err := http.Get(link)
	if err != nil {
		c <- err.Error()
		return
	}
	// c <- link + " with status: " + strconv.Itoa(res.StatusCode) + " at index: " + strconv.Itoa(idx)

	fmt.Println(link, "with status:", res.StatusCode, "at index:", idx)
	c <- link
}
