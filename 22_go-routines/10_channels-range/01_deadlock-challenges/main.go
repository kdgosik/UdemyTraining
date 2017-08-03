package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1 // no body to receive the value
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// What can you do to fix it
