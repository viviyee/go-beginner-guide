package main

import "fmt"

type bill struct {
	name     string
	items    map[string]int
	discount float64
}

func newBill(name string) bill {
	return bill{
		name:     name,
		items:    map[string]int{},
		discount: 0,
	}
}

func main() {
	bill_1 := newBill("naruto")
	fmt.Println(bill_1)
}
