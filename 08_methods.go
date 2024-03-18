package main

import (
	"fmt"
	"strings"
)

/* functions */
func sayName(name string) {
	fmt.Printf("Guten Tag. Ich heisse %s. \n", name)
}

func sayAge(age int) {
	fmt.Printf("Ich bin %d. \n", age)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

/* return */
func bankInterest(amount float64, rate float64, year float64) float64 {
	return amount * rate * year
}

/* return multiples */
func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	words := strings.Split(name, " ")

	var initials []string
	for _, v := range words {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}

func methods() {
	sayName("Anna")
	sayAge(26)

	cycleNames([]string{"Katze", "Hund", "Papagei", "Hamster"}, sayName)

	fmt.Println(bankInterest(100000, 0.08, 1))

	fn1, sn1 := getInitials("Cloud Strife")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("Tifa Lockhart")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Reno")
	fmt.Println(fn3, sn3)
}
