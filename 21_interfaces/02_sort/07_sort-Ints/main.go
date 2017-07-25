package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

}

//https:www.golang.org/pkg/sort/#Sort
//https:www.golang.org/pkg/sort/#Interface
