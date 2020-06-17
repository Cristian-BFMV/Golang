package main

import "fmt"

func main() {
	// & hace referencia a la dirección de memoria y * hace referencia al valor que esta almacenado en esa dirección
	var x int = 1
	var y int
	// *int hace que la variable sea un apuntador a un integer
	var ip *int
	// ip esta apuntando a la dirección de memoria de x
	ip = &x
	// y toma el valor de la dirección de memoria a la cual esta apuntando ip
	y = *ip
	fmt.Println(y)
	// Retorna un apuntador a un integer
	ptr := new(int)
	// Almacena el número 3 a la dirección de memoria a la que apunta ptr
	*ptr = 3
	fmt.Println(*ptr)
}
