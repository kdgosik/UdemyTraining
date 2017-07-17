package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)
	// 3 is the lenght and capacity
	//  length - is the number of elements referred to by the slice
	// capcity - number of elements in the underlying array

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is the length
	// 5 is the capacity

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
