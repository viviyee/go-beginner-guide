package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a bill for: ")
	b := newBill(name)

	// Infinity looop
	for {
		item_name, _ := getInput(reader, "Item name: ")
		if item_name == "" {
			break
		}
		item_price, _ := getInput(reader, "Item price: ")
		p, err := strconv.ParseInt(item_price, 10, 64)

		if err != nil {
			fmt.Println("Invalid input for item price. The Item won't be added to the bill.")
			continue
		} else {
			b.addItem(item_name, int(p))
		}
	}
	discount, _ := getInput(reader, "Enter discount: ")
	if discount != "" {
		d, err := strconv.ParseFloat(discount, 64)
		if err != nil {
			fmt.Println("Invalid input for discount. Discount won't be added to the bill.")
		} else {
			b.updateDiscount(d)
		}
	}

	b.save()
}
