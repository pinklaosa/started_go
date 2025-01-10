package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Details struct {
	name  string
	food  string
	move  string
	sound string
}
type Animals []Details

func (a Animals) InitData() Animals {
	d := append(a, Details{name: "cow", food: "grass", move: "walk", sound: "moo"},
		Details{name: "bird", food: "worms", move: "fly", sound: "peep"},
		Details{name: "snake", food: "mice", move: "slither", sound: "hsss "})
	return d
}

func (a Animals) QueryDetail(detail string, name string) string {

	for _, d := range a {
		if name == d.name {
			switch detail {
			case "food":
				return d.food
			case "move":
				return d.move
			case "sound":
				return d.sound
			default:
				fmt.Print("Error No Data")
			}

		}
	}
	return "No data"
}

func PrintAnimal(detals, name string) {
	var animal Animals
	animal = animal.InitData()
	detail := animal.QueryDetail(detals, name)
	fmt.Println(detail)
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	PrintAnimal("food", "cow")
}
func (c Cow) Move() {
	PrintAnimal("move", "cow")
}
func (c Cow) Speak() {
	PrintAnimal("sound", "cow")
}

type Bird struct{}

func (b Bird) Eat() {
	PrintAnimal("food", "bird")
}
func (b Bird) Move() {
	PrintAnimal("move", "bird")
}
func (b Bird) Speak() {
	PrintAnimal("sound", "bird")
}

type Snake struct{}

func (s Snake) Eat() {
	PrintAnimal("food", "snake")
}
func (s Snake) Move() {
	PrintAnimal("move", "snake")
}
func (s Snake) Speak() {
	PrintAnimal("sound", "snake")
}

func main() {

	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter commands. Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid command. Format: newanimal <name> <type> OR query <name> <info>")
			continue
		}

		command, name, arg := parts[0], parts[1], parts[2]

		switch command {
		case "newanimal":
			var animal Animal
			switch strings.ToLower(arg) {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Unknown animal type. Choose from: cow, bird, snake.")
				continue
			}
			animals[name] = animal
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}
			switch strings.ToLower(arg) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query type. Choose from: eat, move, speak.")
			}

		default:
			fmt.Println("Unknown command. Use 'newanimal' or 'query'.")
		}
	}

}
