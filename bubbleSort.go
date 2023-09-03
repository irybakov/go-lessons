package main

import (
	"fmt"
	"strconv"
)

func swap(nums []int, i int) {
	if nums[i] > nums[i+1] {
		tmp := nums[i]
		nums[i] = nums[i+1]
		nums[i+1] = tmp
	}
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			swap(nums, j)
		}
	}
}

func main() {

	// create slice of 3
	// read the line
	var input string
	var intSlice = make([]int, 0)

	//Promt to enter
	fmt.Println("Hello, enter ints to form the array, then enter S ")
	fmt.Printf("Enter the int number or S: ")

	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	for input != "S" {
		// try to convert to int
		intVar, err := strconv.Atoi(input)
		if err == nil {
			// append to slice
			intSlice = append(intSlice, intVar)
			fmt.Println(intSlice)
		} else {
			fmt.Printf("not an int - %s\n", input)
		}

		// read line
		fmt.Printf("Enter the int number: ")
		fmt.Scanln(&input)
	}

	BubbleSort(intSlice)

	fmt.Println("")
	fmt.Println("Sorted array:")
	fmt.Println(intSlice)

}
