package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Ahmad"
	lastName = "Tsaqif"
	return
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
