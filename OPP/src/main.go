package main

import (
	"fmt"
)

/**
En Go no se tiene el concepto de clase, sin embargo, lo que se hace es crear un tipo de dato
y asociar métodos a ese tipo de dato que se definen unos recievers que indican que esa función
solo puede ser usada por los
*/

type MyInt int

// Definiendo la función asociada a el tipo de dato MyInt con el reciever
func (mi MyInt) Double() int {
	// El reciever es pasado como argumento de forma implicita
	return int(mi * 2)
}

type Person struct {
	name     string
	lastName string
}

//Con struct se puede realizar lo mismo
func (person Person) getFullName() string {
	return person.name + " " + person.lastName
}

/*
	El argumento del reciever es una copia del tipo de dato que recibe, porque es pasado por valor
	por lo que si queremos modificar algo sobre este no podremos hacerlo dado a que es una copia.
	Esto se soluciona pasando el reciever por referencia
*/
func (person *Person) changeName(newName string) {
	person.name = newName
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
	person := Person{name: "Cristian", lastName: "Mendoza"}
	fmt.Println(person.getFullName())
	person.changeName("Camilo")
	fmt.Println(person.getFullName())
}
