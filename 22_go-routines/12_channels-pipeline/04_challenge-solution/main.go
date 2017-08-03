package main

import (
	"fmt"
)

func main() {
	go run_factorial()
}

func run_factorial() {
	for i := 1; i <= 100; i++ {
		for n := range factorial(gen(i)) {
			fmt.Println(i, " - ", n)
		}
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for n := range in {
			for i := n; i > 0; i-- {
				total *= i
			}
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
