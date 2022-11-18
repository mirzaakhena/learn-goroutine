package main

import "fmt"

// new channel
// un-buffer channel
// buffer channel

// send to channel
// receive from channel
// close the channel

func _01() {
	var c1 chan int
	//c1 <- 1
	//<-c1
	//close(c1)
	fmt.Printf("c1 is %T with value %v\n", c1, c1)
}

// c1 is allocated
func _02() {
	var c1 = make(chan int)
	//c1 <- 1
	//<-c1
	//close(c1)
	fmt.Printf("c1 is %T with value %v\n", c1, c1)
}

func _03() {
	var c1 = make(chan int)
	go func(c chan int) {
		<-c
	}(c1)
	c1 <- 1
	c1 <- 2
	fmt.Printf("c1 is %T with value %v\n", c1, c1)
}

func _04() {
	var c1 = make(chan int, 2)
	go func(c chan int) {
		<-c
	}(c1)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	fmt.Printf("c1 is %T with value %v\n", c1, c1)
}

func main() {
	_04()
}
