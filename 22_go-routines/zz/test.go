package main

import "fmt"

func main() {
	c := 10
	for n := range c {
		fmt.Println(n)
	}
}
