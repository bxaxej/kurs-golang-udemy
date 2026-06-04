package main

import "fmt"

// var może być zmieniona w programie, const nie może być zmieniona
var c, python, java bool

var k, j = 1, 2.3

const d int = 42

func main() {

	var i int
	fmt.Println(i, c, python, java, d)

	var css, ruby, html = true, false, "no!"
	fmt.Println(k, j, css, ruby, html)
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t", k, j, css, ruby, html)
}
