package main

import (
	"fmt"
	"math"
)

// Declarando una variable con el tipo de func
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

// Una función puede ser pasada como parametro y ejecutarla dentro de otra función
func applyIt(aFunct func(int) int, val int) int {
	return aFunct(val)
}

func argumentFunc(val int) int {
	return val * 45
}

// Declarando una varible como una función anonima
var testFunc = func(x int) {
	fmt.Println(x)
}

// Función que devuelve una función
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	// Se retorna la función dependiendo de los argumentos recibidos
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

// Función con argumentos variables, se usa ... (elipses), se usan como un slice
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	// Cuando una función tiene la palabra clave defer, esta se ejecuta cuando la función que la contiene termina
	defer fmt.Println("Bye")
	fmt.Println("Hello")
	// Asignando la función a la variable
	funcVar = incFn
	output := funcVar(25)
	output2 := applyIt(argumentFunc, 1)
	// Usando una función anonima como argumento
	output3 := applyIt(func(x int) int {
		return x * 12
	}, 1)
	// Recibe una función cuyo origen es 0, 0
	Dist1 := MakeDistOrigin(0, 0)
	// Recibe una función cuyo origen es 2, 2
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(output, output2, output3)
	// El resultado va a ser diferente debido a que las funciones tiene origines diferentes
	fmt.Println(Dist1(1, 2), Dist2(1, 2))
	// Invocando la función getMax pasandole todos los valores 1 por 1
	fmt.Println(getMax(1, 2, 3, 4))
	// Invocando la función getMax pasandole todos los valores en un slice con el sufijo ...
	valSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(getMax(valSlice...))

}
