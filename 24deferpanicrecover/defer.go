// Defer fn adalah fn yang bisa kita jadwalkan untuk dieksekusi setelah sebuah fn selessai di eksekusi
// Defer fn akan selalu dieksekusi walaupun terjadi error di f yang dieksekusi

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

// Manual //

// func runAppM(value int) {
// 	fmt.Println("Run Application!")
// 	result := 10 / value
// 	fmt.Println(result)
// 	logging()
// }

// Defer //

func runApp() {
	defer logging()
	fmt.Println("Run Application!")
}

func main() {
	// runAppM(0)
	runApp()
}
