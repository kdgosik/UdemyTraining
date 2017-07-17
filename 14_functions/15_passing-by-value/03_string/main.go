package main

import "fmt"

func main() {
	name := "Kirk"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name) // 24
}

func changeMe(z string) {
	fmt.Println(z) //Kirk
	z = "Rocky"
	fmt.Println(z) // Rocky
}
