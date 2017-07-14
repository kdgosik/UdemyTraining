package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) // 0x208...

	var b *int = &a
	fmt.Println(b)  //0x208...
	fmt.Println(*b) // 43

	// b is a pointer;
	// b points to the memory address where an int is stored
	// to see the value in the memory address, and a * in front of
	// this is known as dereferencing
	// the * is an operator in this case
}
