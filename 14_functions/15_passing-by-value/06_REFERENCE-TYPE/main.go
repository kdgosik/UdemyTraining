package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Krik"])
}

func changeMe(z map[string]int) {
	z["Kirk"] = 32
}
