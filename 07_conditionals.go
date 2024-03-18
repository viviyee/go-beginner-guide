package main

import "fmt"

func conditionals() {

	/* Comparison Operators */
	a := 6
	b := 13

	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)

	/* Conditionals */

	// if
	age := 29
	if age >= 20 {
		fmt.Println("You are an adult.")
	}

	// else
	if age >= 20 {
		fmt.Println("You are allowed to drink alcohol.")
	} else {
		fmt.Println("You are NOT allowed to drink.")
	}

	// else if
	temperature := -5
	if temperature > 30 {
		fmt.Println("It's a hot day.")
	} else if temperature < 10 {
		fmt.Println("It's a cold day.")
	} else {
		fmt.Println("It's a lovely day.")
	}

	// break continue
	var numbers []int
	for i := 1; i <= 25; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	for _, num := range numbers {
		if num == 21 {
			break
		}

		if num%7 == 0 {
			fmt.Println("")
			continue
		}
		fmt.Println(num)
	}

	// switch
	char := "C"
	switch char {
	case "A":
		fmt.Println("Ape")
	case "B":
		fmt.Println("Bear")
	case "C":
		fmt.Println("Capybara")
	case "D":
		fmt.Println("Deer")
	case "E":
		fmt.Println("Elephant")
	case "F":
		fmt.Println("Fox")
	default:
		fmt.Println("Default Animal")
	}
}
