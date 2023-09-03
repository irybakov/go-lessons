package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {

	var file string
	var ns = make([]names, 0)

	//Promt to enter name of file
	fmt.Printf("Hello, enter the file name: ")
	fmt.Scanln(&file)

	fs, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(fs)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var txt = strings.Split(scanner.Text(), " ")
		fmt.Println(txt)
		ns = append(ns, names{fname: txt[0], lname: txt[1]})
	}

	fs.Close()

	for idx, name := range ns {
		fmt.Println(idx, "-", name)
	}

}
