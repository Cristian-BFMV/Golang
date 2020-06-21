package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(accelaration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (accelaration*math.Pow(time, 2))/2 + initialVelocity*time + initialDisplacement
	}
}

func main() {
	var accelaration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64
	fmt.Print("Enter a value for acceleration: ")
	fmt.Scan(&accelaration)
	fmt.Print("Enter a value for initial velocity: ")
	fmt.Scan(&initialVelocity)
	fmt.Print("Enter a value for initial displacement: ")
	fmt.Scan(&initialDisplacement)
	fmt.Print("Enter a value for time: ")
	fmt.Scan(&time)
	Displacement := GenDisplaceFn(accelaration, initialVelocity, initialDisplacement)
	fmt.Println(Displacement(time))
}
