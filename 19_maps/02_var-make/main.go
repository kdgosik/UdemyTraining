package main

import "fmt"

func main() {
	myGreeting := make(map[string]string) // took away var and inserted a colon
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
