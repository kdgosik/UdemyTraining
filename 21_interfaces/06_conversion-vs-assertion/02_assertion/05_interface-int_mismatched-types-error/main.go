package main

import (
	"fmt"
)

// invalid operation, due to mismatch type
func main() {
	var val interface{} = 7
	fmt.Println(val + 6)
}
