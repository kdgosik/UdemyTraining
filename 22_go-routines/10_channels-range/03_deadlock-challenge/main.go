package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

// why does this only print 0
// And what can you do to print the numbers 0-9?
