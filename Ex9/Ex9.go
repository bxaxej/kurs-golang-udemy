package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {

	fmt.Printf("KB %d \t %b\n", KB, KB)
	fmt.Printf("MB %d \t %b\n", MB, MB)
	fmt.Printf("GB %d \t %b\n", GB, GB)
	fmt.Printf("TB %d \t %b\n", TB, TB)
	fmt.Printf("PB %d \t %b\n", PB, PB)
	fmt.Printf("EB %d \t %b\n", EB, EB)

}
