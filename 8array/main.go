package main

import "fmt"

func main() {
	var fruits [3]string
	fruits[0] = "Pineapple"
	fruits[1] = "Mango"
	fruits[2] = "Watermelon"

	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
	fmt.Println(fruits)

	values := [3]int {
		90,
		88,
		77,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))
}