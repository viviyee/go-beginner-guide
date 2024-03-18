package main

import "fmt"

func loops() {

	/* While Loop */
	x := 0
	for x < 5 {
		fmt.Println("x ", x)
		x++
	}

	fmt.Println("")

	/* For Loop */
	for y := 10; y < 15; y++ {
		fmt.Println("y ", y)
	}

	fmt.Println("")

	/** Looping Collections */
	// Arrays Slices Maps Channels

	// index looping
	var scandinavian_countires [3]string = [3]string{"Sweden", "Norway", "Denmark"}
	for i := 0; i < len(scandinavian_countires); i++ {
		fmt.Println("Scandinavian : ", scandinavian_countires[i])
	}

	fmt.Println("")

	// range looping (Arrays and Slices)
	cis_countries := []string{"Russia", "Kazakhstan", "Uzbekistsan", "Kyrgyzstan", "Turkmenistan", "Tajikistan", "Moldova", "Belarus", "Armenia", "Azerbaijan"}
	for index, value := range cis_countries {
		fmt.Printf("CIS %d %s \n", index, value)
	}

	fmt.Println("")

	// if you don't want the index
	for _, value := range cis_countries {
		fmt.Println(value)
		value = "new country name" // this doesn't affecct the original value
	}

	fmt.Println(cis_countries)

	// range looping (Maps)
	japan := map[string]string{
		"prime_minister": "Fumio Kishida",
		"capital":        "Tokyo",
		"population":     "123 M",
		"exchange_rate":  "1 USD = 150 JPY",
	}
	for key, value := range japan {
		fmt.Printf("%s: %s \n", key, value)
	}

	fmt.Println("")

	// range looping (Strings)
	for index, char := range "Tokyo" {
		fmt.Printf("%d %v %c \n", index, char, char)
	}
}
