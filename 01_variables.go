package main

import "fmt"

func variables() {
	// Declare a Variable
	var continent string
	continent = "Asia"
	fmt.Println(continent)

	// Declare and Initialize a Variable
	var country string = "Japan"
	fmt.Println(country)

	// Type Inference
	var city = "Tokyo"
	fmt.Println(city)

	// Declare Multiple Variables
	var (
		city_2 string = "Osaka"
		city_3 string = "Yokohoma"
	)
	fmt.Println(city_2, city_3)

	// Short Declaration
	station := "Ikebukuro"
	fmt.Println(station)

	// Short Declaration (Multiple)
	station_2, station_3 := "Kanamecho", "Senkawa"
	fmt.Println(station_2, station_3)
}
