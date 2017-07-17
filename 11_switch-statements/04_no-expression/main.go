package main

import "fmt"

func main() {

	myFriendName := "Mar"

	switch {
	case len(myFriendName) == 2:
		fmt.Println("Wassup my friend of length 2")
	case myFriendName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendName == "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case myFriendName == "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case myFriendName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched, this is the default")
	}
}
