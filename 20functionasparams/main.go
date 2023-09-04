package main

import "fmt"

// func sayHello(name string, filter func(string) string) {
// 	name = filter(name)
// 	fmt.Println("Hello", name)
// }

// Type declaration //
// type Filter func(string, string, string, int, int, bool) string
type Filter func(string) string

func sayHello(name string, filter Filter) {
	name = filter(name)
	fmt.Println("Hello", name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Tsaqif", spamFilter)
	sayHello("Naufal ", spamFilter)
}
