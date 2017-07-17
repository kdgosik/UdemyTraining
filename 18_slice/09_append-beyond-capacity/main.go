package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is the length
	// 5 is the capacity

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zao an")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
