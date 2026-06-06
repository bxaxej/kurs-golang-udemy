package main

import "fmt"

const ( // iota jest licznikiem i domysla sie ze nastepne tez maja byc liczone
	_ = iota
	a
	b
	c
	d
	e
	f
)

func main() {

	fmt.Printf("%d \t %b\n", 1, 1) // %d dziesiętny, %b binarny
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
