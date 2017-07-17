package main

import "fmt"

// switch on type
// -- normally we switch on value or variable
// -- go allows you to switch on type of variable

type Contact struct {
	greeting string
	name     string
}

// we'll learn more about interface later
func SwitchOnType(x interface{}) {

	switch x.(type) { // this is an assert, asserting that "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {

}
