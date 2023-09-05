package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put ur host")
	var user *string = flag.String("user", "root", "Put ur database user")
	var password *string = flag.String("password", "", "Put ur database password")

	flag.Parse()

	fmt.Println(*host, *user, *password)
	// fmt.Println(host, user, password)
}
