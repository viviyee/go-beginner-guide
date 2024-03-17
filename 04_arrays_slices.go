package main

import "fmt"

func main() {

	// Arrays (fixed-sized collection) [n]T
	var languages [4]string = [4]string{"English", "Japanese", "Finnish", "German"}
	fmt.Println(languages, len(languages))

	var letters = [5]byte{'A', 'B', 'C', 'D', 'E'} // type inferred
	fmt.Println(letters, len(letters))

	/* test */
	var test_1 = [5]int{1, 2, 3}
	fmt.Println(test_1) // [1 2 3 0 0] Default 0

	var test_2 = [5]bool{false, false, true}
	fmt.Println(test_2) //[false false true false false] Default false

	var test_3 = [5]string{"int8", "int16", "int32", "int64"}
	fmt.Println(test_3) // Default ""

	// Slices (dynamically-sized array) []T
	var country_codes []string = []string{"BE", "LU", "CH", "AT"}
	var phone_codes = []int{32, 352, 41, 43}
	fmt.Println(country_codes, phone_codes)

	// Append a Slice
	country_codes = append(country_codes, "CZ", "HU") // append func returns a new slice
	fmt.Println(country_codes)

	// Slicing Arrays and Slices (returns slices)
	fmt.Println(languages[1:3])
	fmt.Println(country_codes[2:])
	fmt.Println(letters[:4])
}
