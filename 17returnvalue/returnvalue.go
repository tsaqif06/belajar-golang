package main

import "fmt"

func hello(mas string, adik string) string {
	return "Halo " + mas + ", adik kamu " + adik + " ya!!"
}

func main() {
	hello := hello("Tsaqif", "Naufal")
	fmt.Println(hello)
}
