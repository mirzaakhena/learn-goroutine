package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan<- : only accept input, we can only send to it
// <-chan : only accept output, we can only receive from it

func cakeMaker(kind string, number int, to chan<- string) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		to <- kind
	}
	close(to)
}

func main() {
	chocolateChan := make(chan string, 8)
	redvelvetChan := make(chan string, 8)

	go cakeMaker("chocolate", 15, chocolateChan)
	go cakeMaker("red velvet", 30, redvelvetChan)

	moreChocolate := true
	moreRedvelvet := true

	var cake string

	for moreChocolate || moreRedvelvet {
		select {
		case cake, moreChocolate = <-chocolateChan:
			if moreChocolate {
				fmt.Printf("Got cake #1 %s\n", cake)
			}

		case cake, moreRedvelvet = <-redvelvetChan:
			if moreRedvelvet {
				fmt.Printf("Got cake #2 %s\n", cake)
			}

		case <-time.After(1000 * time.Millisecond):
			fmt.Printf("timeout!\n")
			return
		}

	}

}
