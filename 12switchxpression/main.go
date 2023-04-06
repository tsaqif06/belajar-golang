package main

import "fmt"

func main() {
	role := "ytt"

	// switch role {
	// case "admin":
	// 	fmt.Println("Selamat Datang Admin")
	// case "user":
	// 	fmt.Println("Selamat Datang User")
	// default:
	// 	fmt.Println("Harap Login Terlebih Dahulu")
	// }

	// shorthand

	// switch length := len(role); length == 2 {
	// case true:
	// 	fmt.Println("Mntap")
	// case false:
	// 	fmt.Println("Anjay")
	// }

	// switch tanpa kondisi
	length := len(role)
	switch {
	case length == 2:
		fmt.Println("Hai")
	default:
		fmt.Println("Helo")
	}
}
