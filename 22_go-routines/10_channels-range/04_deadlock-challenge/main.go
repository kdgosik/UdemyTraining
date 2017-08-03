package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
	}
}

// This prints the numbers but we received an error:
// fatal error: all goroutines are asleep - deadlock
// Where is the deadlock?
// Why are all goroutines asleep
// How can we fix this?
