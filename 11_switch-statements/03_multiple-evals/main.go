package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim err or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Wassup Both of you")
		fallthrough
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian/Sushant")
	}
}
