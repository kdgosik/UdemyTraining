package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) // 0x208...

	var b *int = &a
	fmt.Println(b)  //0x208...
	fmt.Println(*b) // 43

	*b = 42        // b says "The value at this address change to 42"
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass)
	// and then we can stil change the value of whatever is stored at the memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around an address

	// everything is pass by value in Go
	// when we pass a memory address, we are passing a value
}
