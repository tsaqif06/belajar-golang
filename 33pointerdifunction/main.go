package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddress(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Malang", "Jawa Timur", ""}
	addressPointer := &address
	ChangeAddress(addressPointer)

	fmt.Println(addressPointer)
}
