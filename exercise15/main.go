package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go func() {

		//c <- "hello"

	}()

	select {
	case msg := <-c:
		fmt.Printf("%s\n", msg)
	}

	time.Sleep(200 * time.Millisecond)
}
