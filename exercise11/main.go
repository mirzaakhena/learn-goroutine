package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "setiap 500ms"
			time.Sleep(500 * time.Millisecond)
		}

	}()

	go func() {

		for {
			c2 <- "setiap 2000ms"
			time.Sleep(2000 * time.Millisecond)
		}

	}()

	for {

		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)

		}
	}

}
