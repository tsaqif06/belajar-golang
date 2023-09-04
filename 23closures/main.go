// Closure adalah kemampuan sebuah fn berinteraksi dgn data2 disekitarnya dalam scope yg podo
package main

import "fmt"

func main() {
	name := "Tsaqif"
	counter := 0

	increment := func() {
		// name = "Naufal" // menimpa value "Tsaqif" jadi "Naufal"
		name := "Naufal" // membuat var baru
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
