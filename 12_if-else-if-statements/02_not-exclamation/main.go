package main

import "fmt"

func main() {
	if !true {
		fmt.Println("this did not run")
	}

	if !false {
		fmt.Println("this ran")
	}
}
