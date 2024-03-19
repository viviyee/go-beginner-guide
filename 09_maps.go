package main

import "fmt"

func maps() {
	/* Initialize a Map [T]T */
	marks := map[string]int{
		"James":    65,
		"William":  68,
		"Benjamin": 95,
		"Daniel":   72,
	}
	fmt.Println(marks["Benjamin"])
	fmt.Println(marks)

	/* Delcare and Initialize a Map [T]T */
	var grades map[string]string
	grades = map[string]string{
		"Emily":     "C",
		"Charlotte": "A",
		"Abigail":   "D",
		"Susan":     "B",
	}
	fmt.Println("Charlotte")
	fmt.Println(grades)

	/* Another way to Initialize a Map */
	scholarships := make(map[uint]int)
	fmt.Println(scholarships)

	/* Checking for a Key */
	abigail_grade, abigail_check := grades["Abigail"]
	fmt.Println(abigail_grade, abigail_check)

	noah_mark, noah_check := marks["Noah"]
	fmt.Println(noah_mark, noah_check) // 0 false (still returns 0)

	/* Iterating over a Map */
	for name, mark := range marks {
		fmt.Println(name, mark)
	}

	/* Adding a Key-Value Pair to a Map */
	marks["Liam"] = 75
	fmt.Println(marks)

	/* Deleting a Key-Value Pair from a Map */
	delete(grades, "Susan")
	fmt.Println(grades)
}
