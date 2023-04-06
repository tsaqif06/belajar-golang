package main

import "fmt"

func main() {
    fruits := map[string]int{
		"apple" : 4,
		"banana" : 5,
		"orange" : 2,
	}

	fruits["watermelon"] = 1

    fmt.Println(fruits)

	people := make(map[string]int)

    people["login"] = 34
    people["logout"] = 12
    people["hacker"] = 2

    fmt.Println(people)
	delete(people, "hacker")
    fmt.Println(people)
}