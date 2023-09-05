package helper

import "fmt"

/** Access Modifier
Jika function atau variable diawali oleh huruf besar
maka dia bisa diakses di package luar,
jika awalannya kecil, maka tidak bisa diakses oleh package luar
*/

func SayHello(name string) {
	fmt.Println("Hello", name)
}
