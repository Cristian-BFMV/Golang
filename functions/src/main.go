package main

import "fmt"

//Función ejecutada por valor, lo que significa que se hace una copia del valor al parametro
func fooByValue(y int) int {
	return y + 1
}

//Función ejecutada por valor, lo que significa apunta a la dirección de memoria de la variable
func fooByReference(y *int) {
	*y = *y + 1
}

// Una función puede tener multiples retornos
func fooMultipleReturn(x int) (int, string) {
	return x + 1, "El resultado es: "
}

// Se le puede pasar a una función un arreglo por referencia
func fooArray(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

// Para manipular arreglos en las funciones es mejor pasar slices que tienen un
func fooSlice(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	x := 2
	x = fooByValue(x)
	fmt.Println(x)
	w := 3
	fooByReference(&w)
	fmt.Println(w)
	z := 4
	a, b := fooMultipleReturn(z)
	fmt.Println(b, a)
	c := [3]int{1, 2, 3}
	fooArray(&c)
	fmt.Println(c)
	//Una forma de llamar a la función con un slice
	d := [4]int{1, 2, 3, 4}
	e := d[0:len(d)]
	fooSlice(e)
	fmt.Println(e)
	//Otra forma de llamar a la función con un slice
	// Cuando no se le especifica el lenght Go lo reconoce como un slice
	f := []int{1, 2, 3, 4}
	fooSlice(f)
	fmt.Println(f)
}
