package main

import "fmt"

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	person1 := Person{name: "Cristian", addr: "Cra 34", phone: "3148945884"}
	fmt.Println(person1)
	fmt.Println(person1.name)
	fmt.Println(person1.addr)
	fmt.Println(person1.phone)
}
