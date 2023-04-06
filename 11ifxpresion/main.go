package main

import "fmt"

func main() {
	role := "ytt"

	// if role == "admin" {
	// 	fmt.Println("Selamat Datang Admin")
	// } else if role == "user" {
	// 	fmt.Println("Selamat Datang User")
	// } else {
	// 	fmt.Println("Harap Login Terlebih Dahulu")
	// }

	// shorthand
	if length := len(role); length == 2 {
		fmt.Println("anjay")
	} else {
		fmt.Println("rbuh")
	}
}
