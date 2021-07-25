package main

import (
"fmt"
"math"
)

func main() {
	var inputAcceleration float64
	fmt.Print("Enter acceleration:")
	fmt.Scanln(&inputAcceleration)
	var inputVelocity float64
	fmt.Print("Enter the initial velocity:")
	fmt.Scanln(&inputVelocity)
	var initialDisplacement float64
	fmt.Print("Enter initial displacement:")
	fmt.Scanln(&initialDisplacement)
	var inputTime float64
	fmt.Print("Enter time:")
	fmt.Scanln(&inputTime)
	fn := GenDisplaceFn(inputAcceleration, inputVelocity, initialDisplacement)
	fmt.Println(fn(inputTime))
	//fmt.Println(fn(5))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5*a*math.Pow(t,2) + v0*t + s0)
	}
}
