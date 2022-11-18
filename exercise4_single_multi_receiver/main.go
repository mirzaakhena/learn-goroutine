package main

import "fmt"

func lucasoid(a, b, n int) int {

	//fmt.Printf("a:%d b:%d n:%d\n", a, b, n)

	if n == 0 {
		return a
	}

	if n == 1 {
		return b
	}

	return lucasoid(a, b, n-1) + lucasoid(a, b, n-2)
}

func multiReceiver() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() { ch1 <- lucasoid(0, 1, 20) }()
	go func() { ch2 <- lucasoid(0, 1, 30) }()
	go func() { ch3 <- lucasoid(0, 1, 40) }()

	select {
	case message := <-ch1:
		fmt.Printf("ch1 %d\n", message)
	case message := <-ch2:
		fmt.Printf("ch2 %d\n", message)
	case message := <-ch3:
		fmt.Printf("ch3 %d\n", message)
	}
}

func singleReceiver() {
	ch := make(chan int)

	go func() { ch <- lucasoid(0, 1, 20) }()
	go func() { ch <- lucasoid(0, 1, 30) }()
	go func() { ch <- lucasoid(0, 1, 40) }()

	fmt.Printf("ch %d\n", <-ch)
}

func main() {

	singleReceiver()

}
