package main

import "fmt"

func main() {
	
	//math

	a := 2
	b := 3
	c := 1

	c += 3

	fmt.Println(a + b - c * b)

	// perbandingan

	w := 1
	e := 1
	fmt.Println(w == e)

	// bool

	s := 86
	g := 90
	
	na := s > 80
	lb := g > 85

	nl := na && lb

	fmt.Println(nl)
}