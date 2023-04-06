package main

import "fmt"

func main() {
	// // for while

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// for dengan stmt
	// for counter := 10; counter >= 1; counter-- {
	// 	fmt.Println(counter)
	// }

	//for for slice or array
	// fruits := []string{"Orange", "Lemon", "Melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i])
	// }

	// for range
	fruits := []string{"Orange", "Lemon", "Melon"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	hp := map[string]int {
		"asus" : 12,
		"lenovo" : 5,
	}
	for key, value := range hp {
		fmt.Println("key", key, "value", value)
	}
}
