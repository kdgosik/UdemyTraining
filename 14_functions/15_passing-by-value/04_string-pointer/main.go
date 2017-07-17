package main

import "fmt"

func main() {

	name := "Kirk"
	fmt.Println(&name) //0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  // 24
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
}
