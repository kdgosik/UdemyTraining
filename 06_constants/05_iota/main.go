package main

import "fmt"

const (
	A = iota //0
	B = iota //1
	C = iota //2
)

const (
	D = iota //0
	E = iota //1
	F = iota //2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
