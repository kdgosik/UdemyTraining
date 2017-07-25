package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning.",
		1: "Bonjour.",
		2: "Beunos dias!",
		3: "Bongiorno.",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)
	fmt.Println(myGreeting)
}
