package main

import (
	"fmt"
	"sync"
)

func Increment(ptr *int, wg *sync.WaitGroup) {
	i := *ptr

	fmt.Printf("i: %d\n", i)

	*ptr = i + 1

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	val := 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go Increment(&val, &wg)
	}

	wg.Wait()

	fmt.Printf("hasil : %d\n", val)

}
