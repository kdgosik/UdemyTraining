package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// loop over words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:122])
	// fmt.Println("*************")
	// for i := 28; i <= 126; i++ {
	// fmt.Printf("%v - %c - %v \n", i, i, buckets[i]
	// }
}

func HashBucket(word string) int {
	return int(word[0])
}
