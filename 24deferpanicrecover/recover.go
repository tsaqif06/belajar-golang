// Recover untuk menangkap data panic, recover akan menghentikan proses panic, sehingga program akan terus berjalan
package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	// message := recover()
	// fmt.Println("Error dengan message :", message)
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("test")
}
