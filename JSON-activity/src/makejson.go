package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter your address")
	fmt.Scan(&address)
	user := map[string]string{"name": name, "address": address}
	barr, _ := json.Marshal(user)
	fmt.Println(string(barr))
}
