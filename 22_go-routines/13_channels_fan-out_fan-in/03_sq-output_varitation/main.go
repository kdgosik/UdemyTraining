package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	//FAN Out
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	//Fan In
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9 or 9 then 4
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

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in cs
	// output copies values from c to out until c is closed, then calling wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out all the output goroutine are
	// done.  This must start after the wg.Add() call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
