package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is the length
	// 5 is the capacity

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])
}
