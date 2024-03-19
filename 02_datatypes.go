package main

import "fmt"

func datatypes() {
	// string
	var name string = "Anna"
	fmt.Println(name)

	// int (positive, negative)
	/* int  int8  int16  int32  int64 */
	var temperature_1 int = 34
	var temperature_2 int = -5
	fmt.Println(temperature_1, temperature_2)

	// bit and byte
	/*
		# Don't specify a size unless you really need it.

		bit = smallest unit of data, either 0 or 1. a binary digit
		8 bits = 1 byte

		int, uint, unitptr
		(depends on platform) 32 bits in 32 bit-systems, 64 bits in 64 bit-systems

		int8 = 8 bits/ 1 byte (-128 to 127)
		int16 = 16 bits/ 2 bytes (-32 k to 32 k)
		int32 = 32 bits/ 4 bytes (-2 billion to 2 billion)
		int64 = 64 bits/ 8 bytes (-9.2 sextillion to 9.2 sextillion) 1 sextillion = 10 pow 21

		uint8 = 0 to 255
		uint16 = 0 to 65 k
		uint32 = 0 to 4 G (1 G = 10 pow 9)
		uint64 = 0 to 18 E (1 E = 10 pow 18)
	*/

	// uint (unsigned integer, only positive)
	/* uint uint8 uint16 uint32 uint64 uintptr */
	var age uint = 30
	fmt.Println(age)

	// float32, float64
	var height float32 = 150.5
	var weight float64 = 43.3
	/* use float64 more, it has a slightly higher precision */
	gpa := 3.4 // inferred as float64
	fmt.Println(height, weight, gpa)

	// byte // alias for uint8
	/* represents ASCII characters */
	var num byte = 255
	var char byte = 'B'
	fmt.Println(num, char)

	// rune // alias for int32
	/* represents a Unicode code point */
	var currency rune = '€'
	fmt.Printf("Unicode code point for Euro sign: %d\n", currency)
	fmt.Printf("Character representation: %c\n", currency)

	// complex64, complex128
	/* complex numbers with real and imaginary parts */
	var imaginary_1 complex64 = 2 + 3i
	imaginary_2 := 3 + 4i // inferred as complex128
	fmt.Println(imaginary_1, imaginary_2)

	// bool (true or false)
	var isAdult bool = true
	var liveInUSA bool = false
	fmt.Println(isAdult, liveInUSA)

	// Array and Slice

	// Map

	// Pointer

	// Struct

	// Interface

	/*
		Group A (Pass by value) (Non Pointer Values)
		strings
		ints
		floats
		booleans
		arrays
		structs

		Group B (pass by pointer) (Pointer Wrapper Values)
		slices
		maps
		functions
	*/
}
