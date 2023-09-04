package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) sayHello(name string) {
	fmt.Println("Hello "+name+" My Name is", c.Name)
}

func main() {
	var tsaqif Customer
	tsaqif.Name = "Tsaqif"
	tsaqif.Address = "Malang"
	tsaqif.Age = 17

	// Struct Literals

	naufal := Customer{
		Name:    "Naufal",
		Address: "Malang",
		Age:     12,
	}

	zaki := Customer{"Zaki", "Malang", 10}

	// End Struct Literals

	fmt.Println(tsaqif)
	tsaqif.sayHello(naufal.Name)
	fmt.Println(naufal)
	fmt.Println(zaki)
	// fmt.Println(tsaqif.Name)
	// fmt.Println(tsaqif.Address)
	// fmt.Println(tsaqif.Age)
}
