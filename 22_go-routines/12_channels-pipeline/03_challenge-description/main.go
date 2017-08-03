package main

import (
	"fmt"
)

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
Challenge #1
-- change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the pipeline pattern to accomplish this

Post Discussion
-- What realizations did you have with working with concurrent patterns
-- ...


*/
