package main

import (
	"fmt"
	"net/http"
	"sync"
)

// chan<- : only accept input, we can only send to it
// <-chan : only accept output, we can only receive from it

func webGetWorker(in <-chan string, wg *sync.WaitGroup) {
	for {
		url := <-in
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Printf("GET %s: %d\n", url, res.StatusCode)
		wg.Done()
	}
}

// single sender many receiver

func main() {

	work := make(chan string)
	var wg sync.WaitGroup
	numWorker := 10

	for i := 0; i < numWorker; i++ {
		go webGetWorker(work, &wg)
	}

	urls := [6]string{
		"http://notexists.com",
		"http://twitter.com",
		"http://go.dev",
		"http://detik.com",
		"http://paxelmarket.co",
		"http://blibli.com",
	}

	for i := 0; i < 100; i++ {
		for _, url := range urls {
			wg.Add(1)
			work <- url
		}
	}

	wg.Wait()

}
