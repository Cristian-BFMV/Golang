package main

import (
	//Importando el paquete de JSON
	"encoding/json"
	"fmt"
)

//Para que pueda funcionar el paquete de json es necesario tener los nombres del struct en mayusculas
type Person struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

func main() {
	var p2 Person
	p1 := Person{Name: "Cristian", Addr: "Street 44", Phone: "31236"}
	//Convierte el struct en una representación([]byte) de un JSON
	barr, err := json.Marshal(p1)
	if err == nil {
		fmt.Println(string(barr))
		//Convierte la representación([]byte) de un JSON a un struct
		err1 := json.Unmarshal(barr, &p2)
		if err1 == nil {
			fmt.Println(p2)
		}
	}
}
