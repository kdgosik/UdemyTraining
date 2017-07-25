package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
