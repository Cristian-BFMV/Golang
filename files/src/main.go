package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("src/text.txt")
	if err == nil {
		data := make([]byte, 100)
		file.Read(data)
		for _, v := range data {
			if v != 0 {
				fmt.Println(string(v))
			}
		}
		file.Close()
	} else {
		fmt.Println(err)
	}
}
