package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// create slice of 3
	// read the line
	var input string
	var intSlice = make([]int, 3)

	//Promt to enter
	fmt.Printf("Hello, enter the int number: ")

	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	for input != "X" {
		// try to convert to int
		intVar, err := strconv.Atoi(input)
		if err == nil {
			// append to slice
			intSlice = append(intSlice, intVar)
			sort.Ints(intSlice)
			fmt.Printf("slice, %s!\n", intSlice)
		} else {
			fmt.Printf("not an int - %s\n", input)
		}

		// read line
		fmt.Printf("Enter the int number: ")
		fmt.Scanln(&input)
	}

}
