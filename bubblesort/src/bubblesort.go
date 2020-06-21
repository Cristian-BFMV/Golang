package main

import "fmt"

func BubbleSort(slice []int) {
	sliceLength := len(slice)
	for range slice {
		for j := 0; j < sliceLength-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func main() {
	slice := make([]int, 0, 10)
	var userInput int
	for i := 0; i < 10; i++ {
		fmt.Println("Enter a number")
		fmt.Scan(&userInput)
		slice = append(slice, userInput)
	}
	BubbleSort(slice)
	fmt.Println(slice)
}
