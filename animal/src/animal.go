package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println("Animal food: " + animal.food)
}

func (animal Animal) Move() {
	fmt.Println("Animal movement type: " + animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println("Animal speak tpye: " + animal.noise)
}

var cow Animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
var bird Animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
var snake Animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

func main() {
	var animalInput string
	var actionInput string
	ban := true
	for ban {
		fmt.Println(">")
		fmt.Scan(&animalInput, &actionInput)
		PrintAction(animalInput, actionInput)
	}
}

func PrintAction(animal, action string) {
	animal = strings.ToLower(animal)
	action = strings.ToLower(action)
	switch animal {
	case "cow":
		switch action {
		case "eat":
			cow.Eat()
		case "move":
			cow.Move()
		case "speak":
			cow.Speak()
		}
	case "bird":
		switch action {
		case "eat":
			bird.Eat()
		case "move":
			bird.Move()
		case "speak":
			bird.Speak()
		}
	case "snake":
		switch action {
		case "eat":
			snake.Eat()
		case "move":
			snake.Move()
		case "speak":
			snake.Speak()
		default:
			fmt.Println("Wrong action")
		}
	default:
		fmt.Println("Wrong animal")
	}
}
