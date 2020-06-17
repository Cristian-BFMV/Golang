package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	var ban bool = false
	fmt.Println("Enter a word")
	fmt.Scan(&word)
	word = strings.ToLower(word)
	if strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") {
		if strings.Index(word, "a") != -1 {
			ban = true
		}
	}
	if ban {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
