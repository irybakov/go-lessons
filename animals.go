package main

import (
	"fmt"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}

func (c *Cow) Eat() string   { return "grass" }
func (c *Cow) Move() string  { return "walk" }
func (c *Cow) Speak() string { return "moo" }

type Bird struct{}

func (b *Bird) Eat() string   { return "worms" }
func (b *Bird) Move() string  { return "fly" }
func (b *Bird) Speak() string { return "peep" }

type Snake struct{}

func (s *Snake) Eat() string   { return "mice" }
func (s *Snake) Move() string  { return "slither" }
func (s *Snake) Speak() string { return "hsss" }

func CreateAnimal(anmls map[string]Animal, name string, atype string) {
	switch atype {
	case "cow":
		anmls[name] = new(Cow)
	case "bird":
		anmls[name] = new(Bird)
	case "snake":
		anmls[name] = new(Snake)
	}
}

func QueryAnimal(anmls map[string]Animal, name string, info string) {
	switch info {
	case "eat":
		fmt.Printf("%s", anmls[name].Eat())
	case "move":
		fmt.Printf("%s", anmls[name].Move())
	case "speak":
		fmt.Printf("%s", anmls[name].Speak())
	}
}

func ProcessRequest(anmls map[string]Animal, cmd string, param1 string, param2 string) {
	switch cmd {
	case "newanimal":
		CreateAnimal(anmls, param1, param2)
	case "query":
		QueryAnimal(anmls, param1, param2)
	}
}

func main() {

	var anmls map[string]Animal = make(map[string]Animal)

	for true {
		var cmd, name, request string
		fmt.Printf(">")
		_, err := fmt.Scanln(&cmd, &name, &request)
		if err != nil {
			fmt.Printf("¯\\_(ツ)_/¯ - %s\n", err)
		}

		fmt.Println("---")
		ProcessRequest(anmls, cmd, name, request)
		fmt.Println("")
	}

}
