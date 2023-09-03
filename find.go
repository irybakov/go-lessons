package main

import (
	"fmt"
	"strings"
)

func main() {

	var found bool = true
	var input string

	//Promt to enter
	fmt.Printf("Enter a word:\n")
	_, err := fmt.Scan(&input)

	if err != nil {
		return
	}
	// find letters 'i'
	var idx int
	idx = strings.Index(input, "i")
	found = found && (idx > -1)

	// find letters 'a'
	idx = strings.Index(input, "a")
	found = found && (idx > -1)

	// find letters 'i' & 'a' & 'n'
	idx = strings.Index(input, "n")
	found = found && (idx > -1)

	//found would be true if 'i' & 'a' & 'n' were found
	//print the result
	if found {
		fmt.Printf("Found")
	} else {
		fmt.Printf("Not Found")
	}

}
