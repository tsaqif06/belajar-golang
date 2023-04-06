package main

import "fmt"

func main() {
    months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",}

	slice1 := months[4:7]
    fmt.Println(slice1)
    fmt.Println(len(slice1))
    fmt.Println(cap(slice1))

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Apa Ya")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Apa Ya"
	newSlice[1] = "APApa Ya"
	fmt.Println("newSlice = ", newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
}