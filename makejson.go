package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string

	//Promt to enter name
	fmt.Printf("Hello, enter the Name: ")
	fmt.Scanln(&name)

	//Promt to enter address
	fmt.Printf("enter the Address: ")
	fmt.Scanln(&address)

	var hmap = map[string]string{"name": name, "address": address}

	json, err := json.Marshal(hmap)
	if err != nil {
		return
	}

	fmt.Printf("json: %s", json)

}
