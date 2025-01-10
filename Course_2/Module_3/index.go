package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name  string
	food  string
	move  string
	sound string
}

type Animals []Animal

func (a Animals) InitMe() Animals {
	d := append(a, Animal{name: "cow", food: "grass", move: "walk", sound: "moo"},
		Animal{name: "bird", food: "worms", move: "fly", sound: "peep"},
		Animal{name: "snake", food: "mice", move: "slither", sound: "hsss "})
	return d
}

func (a Animals) Eat(name string) {
	one, err := a.SearchName(name)
	if err == false {
		PrintErr()
		return
	}
	fmt.Println(one.name+" eat: ", one.food)
}

func (a Animals) Move(name string) {
	one, err := a.SearchName(name)
	if !err {
		PrintErr()
		return
	}
	fmt.Print(one.name + " move: " + one.move)
}

func (a Animals) Speak(name string) {
	one, err := a.SearchName(name)
	if !err {
		PrintErr()
		return
	}
	fmt.Print(one.name + " speak: " + one.sound)
}

func PrintErr() {
	fmt.Println("No data")
}

func (a Animals) SearchName(name string) (Animal, bool) {
	for _, one := range a {
		if one.name == name {
			return one, true
		}
	}
	return Animal{}, false
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	parts := strings.Fields(input)

	if len(parts) != 2 {
		fmt.Println("type name and func. pls")
		return
	}

	var animals Animals
	animals = animals.InitMe()

	switch parts[1] {
	case "eat":
		animals.Eat(parts[0])
	case "move":
		animals.Move(parts[0])
	case "speak":
		animals.Speak(parts[0])
	default:
		PrintErr()
		return
	}

	return

}
