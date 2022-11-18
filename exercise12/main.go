package main

import "fmt"

func main() {

}

func fibWithMap() {
	r := make(map[int]int)
	fib1(10, r)
	fmt.Printf("hasil :%d\n", r[10])
}

func fibWithoutMap() {
	fmt.Printf("hasil :%d\n", fib2(10))
}

func fib1(n int, result map[int]int) {

	fmt.Printf(">>>> %d\n", n)
	_, exist := result[n]
	if exist {
		return
	}

	if n <= 1 {
		result[n] = 1
		return
	}

	fib1(n-1, result)
	fib1(n-2, result)

	result[n] = result[n-1] + result[n-2]

	return

}

func fib2(n int) int {
	fmt.Printf(">>>> %d\n", n)
	if n <= 1 {
		return 1
	}
	return fib2(n-1) + fib2(n-2)
}
