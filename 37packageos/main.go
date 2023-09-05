package main

import (
	"belajar-golang/36packageinit/database"
	// _ "belajar-golang/36packageinit/database" // "-" adalah blank identifier
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
