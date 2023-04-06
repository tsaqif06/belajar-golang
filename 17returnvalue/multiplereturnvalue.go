package main

import "fmt"

func getFullName() (string, string) {
	return "Ahmad", "Tsaqif"
}

func main() {
	firstName, lastName := getFullName()
	// firstName, _ := getFullName()
	fmt.Println(firstName, lastName)
}
