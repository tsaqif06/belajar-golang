package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	sum := sumAll(10, 10, 10, 10)
	fmt.Println(sum)

	// slice
	slice := []int{1, 2, 3, 4, 5}
	total := sumAll(slice...)
	fmt.Println(total)
}
