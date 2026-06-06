package main

import "fmt"

const Pi = 3.14

func main() {
	v := 42
	s := "42"
	c := 0.34 + 1.09i

	fmt.Printf("The value: %v, of type %T\n", v, v)
	fmt.Printf("The value: %v, of type %T\n", s, s)
	fmt.Printf("The value: %v, of type %T\n", c, c)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
