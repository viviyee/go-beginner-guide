package main

import (
	"fmt"
	"sort"
	"strings"
)

func standard_library() {

	// strings
	sentence := "I am going to Japan next month"

	fmt.Println(strings.Contains(sentence, "Japan"))

	fmt.Println(strings.ReplaceAll(sentence, "a", "@"))

	fmt.Println(strings.ToUpper(sentence))
	fmt.Println(strings.ToLower(sentence))
	fmt.Println(strings.ToTitle(sentence))

	fmt.Println(strings.Index(sentence, "Japan"))
	fmt.Println(strings.Index(sentence, "Singapore")) // -1 if not found

	fmt.Println(strings.Split(sentence, " ")) // returns a slice []string

	// sort
	numbers := []int{6, 1, 13, 0, 100, -6, 4}
	sort.Ints(numbers)
	fmt.Println(numbers) // modify the original slice

	fmt.Println(sort.SearchInts(numbers, 13))
	fmt.Println(sort.SearchInts(numbers, 2)) // returns 3, becase it would fit at index 3, between 1 and 4

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	sort.Strings(planets)
	fmt.Println(planets)

	fmt.Println(sort.SearchStrings(planets, "Jupiter"))
	fmt.Println(sort.SearchStrings(planets, "Pluto")) // returns 5, because it would fit at index 5, between Neptune and Saturn
}
