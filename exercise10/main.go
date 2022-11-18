package main

import (
	"fmt"
	"time"
)

// belajar buffered channel

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	//c <- "three"

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// ================================================================

type Kantong struct {
	Counter int
	Name    string
}

func _01() {

	c := make(chan Kantong)

	go count("kucing", c)

	for message := range c {
		//message, isStillOpen := <-c
		//if !isStillOpen {
		//	break
		//}
		fmt.Printf("%v\n", message)

	}

}

func count(name string, kantongChan chan<- Kantong) {
	for i := 0; i < 5; i++ {
		kantongChan <- Kantong{i, name}
		time.Sleep(500 * time.Millisecond)
	}
	close(kantongChan)
}
