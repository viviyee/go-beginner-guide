package main

import "fmt"

func updateName(name string) {
	name = "Ruyi"
	fmt.Println("updateName", name)
}

func updateNameWithPointer(name *string) {
	*name = "Ruyi"
	fmt.Println("updateNameWithPointer", *name)
}

func pointers() {

	/*
		&variable (copy memory address, get pointer)
		*pointer_variable (get pointer value)
	*/

	color := "white"

	his_fav_color := color // copy value (white)
	color = "black"
	fmt.Println(his_fav_color, color) // white, black
	// his_fav_color won't change, still white

	my_fav_color := &color // copy memory address (0xc000028070)
	color = "dark blue"
	fmt.Println(*my_fav_color, color) // both (dark blue)
	// my_fav_color changed!

	/*
		variable	color		his_fav_color		my_fav_color
		value		white		white				0xc000 (white)
		pointer		0xc000

					dark blue	white				0xc000 (dark blue)
	*/

	/*
		# Confusing part

		use & to get pointer value
		e.g		my_fav_color := &color

		type of my_fav_color is shown with *
		e.g		*string

		use * to get value of a pointer
		e.g		*my_fav_color
	*/

	name := "Anna"
	updateName(name)
	fmt.Println(name)

	updateNameWithPointer(&name)
	fmt.Println(name)
}
