package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a*(t*t))/2 + v0*t + s0
	}
}

func main() {

	// create slice of 3
	// read the line
	var a float64
	var v0 float64
	var s0 float64
	var t float64

	//Promt to enter

	fmt.Printf("Please, enter acceleration a: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return
	}

	fmt.Printf("Please, initial velocity v0: ")
	fmt.Scanln(&v0)

	fmt.Printf("Please, initial displacement s0: ")
	fmt.Scanln(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Please,  time t: ")
	fmt.Scanln(&t)

	fmt.Println(fn(t))

}
