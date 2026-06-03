package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	a, b, c, d, e := 0, 1, 2, 3, 4
	fmt.Printf("%v \t %b \t %X\n", a, a, a)
	fmt.Printf("%v \t %b \t %X\n", b, b, b)
	fmt.Printf("%v \t %b \t %X\n", c, c, c)
	fmt.Printf("%v \t %b \t %X\n", d, d, d)
	fmt.Printf("%v \t %b \t %X", e, e, e)

	y := 42
	z := 42.0

	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742

	fmt.Printf("%v type of %T \n", m, m)

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

	fmt.Println("The time is", time.Now())

	fmt.Println("My favourite numer is", rand.Intn(10))

	fmt.Printf("Now you have %g problems \n", math.Sqrt(49))

	fmt.Printf("Pi is equal to %g", math.Pi)
}
