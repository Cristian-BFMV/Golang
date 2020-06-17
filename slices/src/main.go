package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	sli := []int{1, 2, 3, 4}
	//Inicializa el slice para que tenga como longitud 10
	var sli2 = make([]int, 10)
	//Inicializa el slice para que tenga como longitud 10 pero tiene como capacidad 15
	var sli3 = make([]int, 10, 15)
	//Inicializa el slice para que tenga como longitud 0 pero tiene como capacidad 3
	var sli4 = make([]int, 0, 3)
	// La función append agrega el valor al final incrementando la longitud del slice
	sli4 = append(sli4, 100)
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2), sli, sli2, sli3, sli4)
}
