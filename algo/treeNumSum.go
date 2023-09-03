package main

import (
	"fmt"
	"sort"
)

func main() {

	var input = []int{12, 3, 1, 2, -6, 5, -8, 6}
	var target int = 0
	var result = [][]int{}
	if len(input) < 3 {
		fmt.Println(result)
		return
	}
	sort.Ints(input)
	fmt.Println(input)

	for i := 0; i <= len(input)-2; i++ {
		var left = i + 1
		var right = len(input) - 1
		for left < right {
			var current = input[i] + input[left] + input[right]
			switch {
			case current == target:
				var triplet = []int{input[i], input[left], input[right]}
				result = append(result, triplet)
				left++
				right--
			case current > target:
				right--
			default:
				left++
			}

		}
	}

	fmt.Println(result)

}
