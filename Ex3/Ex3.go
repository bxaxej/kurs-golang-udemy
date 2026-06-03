package main

import "fmt"

// Zwracamy jedną wartość dlatego sam string
func addInt(x, y int) int {
	return x + y
}

func addString(x, y string) string {
	return x + y
}

// Zwracamy dwie wartości dlatego nawias z dwoma stringami
func swap(x, y string) (string, string) {
	return y, x
}

// sum jest wejściem a x, y wyjściem dlatego dwa nawiasy
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(addInt(32, 8))
	fmt.Println(addString("huba ", "buba"))
	sayHello()

	a := addInt(60, 7)
	fmt.Printf("Function says its %v \n", a)

	f, s := swap("hello", "world")
	fmt.Printf("First word is '%s', and the second is '%s' after func swap \n", f, s)

	//tak uzywac obu zmiennych x i y z funkcji
	b, c := split(36)
	fmt.Printf("Splitted results x = '%v' and y = '%v' \n", b, c)

	//tak uzywac tylko jednej bez błędu kompilatora
	_, d := split(36)
	fmt.Printf("Splitted results only for y = '%v' \n", d)

}

func sayHello() {
	fmt.Println("Hello")
}
