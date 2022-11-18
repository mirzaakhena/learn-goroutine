package main

import (
	"fmt"
	"sync"
)

func PrintEven(x int, wg *sync.WaitGroup) {
	if x%2 == 0 {
		fmt.Printf("%d\n", x)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go PrintEven(i, &wg)
	}

	wg.Wait()

}
