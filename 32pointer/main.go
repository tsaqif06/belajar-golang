package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address2 *Address = &address1
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}  // hanya mengubah address2
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // mengubah addres1 dan 2, address 1 seperti ditarik ke milik adress2, dan adress selainnya

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// var address4 *Address = new(Address)
	address4 := new(Address)
	address4.City = "London"
	fmt.Println(address4)
}
