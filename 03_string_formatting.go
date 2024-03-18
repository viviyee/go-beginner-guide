package main

import "fmt"

func string_formatting() {

	// Print
	fmt.Print("Hello ")
	fmt.Print("World.")
	fmt.Print("\nHello Universe.\n")

	// Println
	fmt.Println("Guten Tag.")
	fmt.Println("Auf Wiedersehen.")

	// Printf
	name := "Anna"
	age := 29
	fmt.Printf("My name is %v and I am %v years old.\n", name, age)

	/*
		# Formatting Verbs

		%v 		any values
		%s 		strings
		%d		integers (base 10)
		%f		floating point numbers
		%t		booleans
		%T 		type of variable
		%q		add quotes on string
	*/

	fmt.Printf("My name is %T and I am %T years old.\n", name, age)
	fmt.Printf("My name is %q and I am %q years old.\n", name, age) // quote age as '\x1d'
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	gpa := 3.66
	liveInUSA := false
	fmt.Printf("GPA: %f, Live in USA: %t\n", gpa, liveInUSA)

	fmt.Printf("GPA: %0.1f\n", gpa) // round to 1 decimal place -> 3.7
	fmt.Printf("Age: %05d\n", age)  // pad with zeros to width 5 -> 00029

	// Sprintf
	info := fmt.Sprintf("Name: %s, Age: %d, GPA: %0.1f, Live in USA: %t", name, age, gpa, liveInUSA)
	fmt.Println(info)
}
