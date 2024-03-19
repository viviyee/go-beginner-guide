package main

import (
	"fmt"
	"os"
)

/* Defining a Struct (type struct) */
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

/*
Receiver Methods for Structs
(it lets the function associates with a type)
func (receiver) func_name output_type
*/
func (b *bill) format() string {
	content := fmt.Sprintf("%-20s %s \n\n", "Bill for", "Mr/Ms "+b.name)
	sum := 0

	for name, price := range b.items {
		content += fmt.Sprintf("%-20s %d \n", name, price)
		sum += price
	}
	content += fmt.Sprintf("\n%-20s %d \n", "Sum", sum)

	discount := float64(sum) * b.discount
	total := float64(sum) - discount
	content += fmt.Sprintf("\n%-20s %0.1f \n", "Discount", discount)
	content += fmt.Sprintf("\n%-20s %0.1f \n", "Total", total)

	return content
}

/* Reciever Methods with Pointers */
func (b *bill) updateDiscount(discount float64) {
	b.discount = discount
}

func (b *bill) addItem(name string, price int) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("file was saved")
	}
}

func structs() {
	/* Creating an Instance of Struct */
	bill_1 := newBill("naruto")
	bill_1.items = map[string]int{
		"sushi":     1500,
		"soba":      880,
		"miso soup": 320,
	}
	fmt.Println(bill_1)

	/* Accessing and Modifying Fields */
	fmt.Println(bill_1.name)
	bill_1.name = "sakura"
	fmt.Println(bill_1)

	/* Struct is value-type */
	bill_2 := bill_1
	bill_1.name = "hinata"
	fmt.Println(bill_1, bill_2)
	// bill_1 is hinata, bill_2 is still sakura, doesn't change

	bill_3 := newBill("kiba")
	bill_3.addItem("tonkatsu", 2200)
	bill_3.updateDiscount(0.5)

	bill_1.updateDiscount(0.1)
	bill_1.addItem("tofu", 500)
	bill_1.addItem("udon", 1250)
	bill_1.addItem("natto", 250)

	fmt.Println(bill_1.format())
	fmt.Println(bill_3.format())
}
