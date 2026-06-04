package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // przesunięcie bitowe. Zapis 1<<64 oznacza: 2 do potęgi 64
	z      complex128 = cmplx.Sqrt(-5 + 12i) // complex - liczby zespolone, cmplx - do liczb ujemnych i zespolonych
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var (
		i int
		f float64
		b bool
		s string

		p string = "Huba\tBuba"

		x, y int     = 3, 4
		c    float64 = math.Sqrt(float64(x*x + y*y))
		z    uint    = uint(c)
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s) // %q to bezpieczny quote, czyli da wszystko co jest w zmiennej

	fmt.Printf("Z 'q' %q \t, Z 'v' %v\n", p, p)

	fmt.Println(x, y, z)

}
