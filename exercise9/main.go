package main

import (
	"fmt"
	"sync"
	"time"
)

type OrderID int

type Order struct {
	ID     OrderID
	Weight int
	Cost   int
	Err    error
}

func NewOrder(id, weight int) *Order {
	return &Order{
		ID:     OrderID(id),
		Weight: weight,
		Cost:   0,
		Err:    nil,
	}
}

func main() {

	orders := []*Order{
		NewOrder(1, 40),
		NewOrder(2, 39),
		NewOrder(3, 26),
		NewOrder(4, 32),
		NewOrder(5, 11),
		NewOrder(6, 18),
		NewOrder(7, 21),
		NewOrder(8, 41),
		NewOrder(9, 37),
		NewOrder(10, 16),
		NewOrder(11, 85),
	}

	startTime := time.Now()

	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(1)

		go func(o *Order, wg *sync.WaitGroup) {

			defer wg.Done()

			x, err := doSomeCalculation(o.Weight)
			if err != nil {
				o.Err = err
				fmt.Printf("err : %v, time : %vms\n", err.Error(), time.Since(startTime).Milliseconds())
				return
			}

			o.Cost = x
			fmt.Printf("id : %v, cost : %v, time : %vms\n", o.ID, o.Cost, time.Since(startTime).Milliseconds())

		}(order, &wg)

	}

	wg.Wait()

	fmt.Println("=========")

	hasError := false

	for _, o := range orders {

		if o.Err != nil {
			fmt.Printf("id : %v, weight : %v, cannot calculate cost : %v\n", o.ID, o.Weight, o.Err)
			hasError = true
			continue
		}

		fmt.Printf("id : %v, weight : %v, cost : %v\n", o.ID, o.Weight, o.Cost)
	}

	fmt.Printf("Total time : %vms. Has Error %t\n", time.Since(startTime).Milliseconds(), hasError)

}

func doSomeCalculation(weight int) (int, error) {

	time.Sleep(time.Duration(weight*100) * time.Millisecond)

	if weight%3 == 0 && weight%2 != 0 {
		return 0, fmt.Errorf("gak boleh habis dibagi 3 dan ganjil")
	}

	return 2 * weight, nil
}
