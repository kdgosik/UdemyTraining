package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
Challenge #1:
--Use goroutines and channels to calculate factorials

Challenge #2:
-- Why might you want to use goroutines and channels to calculate factorials?
---- Formulate your own answer
---- Read a few other responses