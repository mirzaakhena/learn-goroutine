package main

import (
	"fmt"
	"time"
)

type Worker struct {
	out chan int
	in  chan int
}

func NewWorker(count int, task func(in <-chan int, out chan<- int)) Worker {

	w := Worker{
		in:  make(chan int),
		out: make(chan int),
	}

	// Create count `multiplyByTwo` goroutines.
	for i := 0; i < count; i++ {
		go task(w.in, w.out)
	}

	return w
}

func main() {

	w := NewWorker(3, func(in <-chan int, out chan<- int) {
		fmt.Println("Initializing goroutine...")
		for {
			num := <-in
			result := num * 2
			time.Sleep(2000 * time.Millisecond)
			out <- result
		}
	})

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data, we can send this in another goroutine
	go func() {
		w.in <- 1
		w.in <- 2
		w.in <- 3
		w.in <- 4
	}()

	// Now we wait for each result to come in
	fmt.Println(<-w.out)
	fmt.Println(<-w.out)
	fmt.Println(<-w.out)
	fmt.Println(<-w.out)

}
