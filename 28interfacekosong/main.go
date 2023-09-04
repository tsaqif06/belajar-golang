package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// var data int = Ups(1)
	var data interface{} = Ups(3)
	// data := Ups(1)
	fmt.Println(data)
}
