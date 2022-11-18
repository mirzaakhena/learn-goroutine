package main

import (
	"fmt"
	"sync"
	"time"
)

// chan<- : only accept input, we can only send to it
// <-chan : only accept output, we can only receive from it

func ManyWriterSingleReader() {

	// we prepare a channel here
	c := make(chan string)

	// we do loop 5 time
	for i := 1; i <= 5; i++ {

		// for each one, we will create goroutine here
		// receive an int and a channel that only receive input
		go func(i int, co chan<- string) {

			// we do loop again here
			for j := 1; j <= 5; j++ {

				// sending the result into channel
				co <- fmt.Sprintf("hi from %d.%d", i, j)
			}

		}(i, c)

	}

	// we do loop again for receive the result
	for i := 1; i <= 25; i++ {
		fmt.Println(<-c)
	}
}

func SingleWriterMultipleReader() {
	c := make(chan int)
	var w sync.WaitGroup
	w.Add(5)

	for i := 1; i <= 5; i++ {
		go func(i int, ci <-chan int) {
			j := 1
			for v := range ci {
				time.Sleep(time.Millisecond)
				fmt.Printf("%d.%d got %d\n", i, j, v)
				j += 1
			}
			w.Done()
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
	w.Wait()
}

func main() {

	ManyWriterSingleReader()
	SingleWriterMultipleReader()

}
