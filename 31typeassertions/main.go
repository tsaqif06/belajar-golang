// Type Assertion = kemampuan merubah tipe data menjadi tipe data yg dikepingini

package main

import "fmt"

func random() interface{} {
	return "test"
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(result)
	// fmt.Println(resultString)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown type")
	}
}
