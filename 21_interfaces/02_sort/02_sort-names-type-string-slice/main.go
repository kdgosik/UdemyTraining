package main

import (
	"fmt"
	"sort"
)

// these are redundant; already inherent in calling StringSlice()

//type StringSlice []string

//func (p StringSlice) Len() int           { return len(p) }
//func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

//https:www.golang.org/pkg/sort/#Sort
//https:www.golang.org/pkg/sort/#Interface
