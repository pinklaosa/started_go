package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	type Data struct {
		fname string
		lname string
	}

	var nameFile string

	var list []Data

	fmt.Print("Enter file name that you need import(.txt): ")
	fmt.Scan(&nameFile)

	dat, _ := os.ReadFile(nameFile + ".txt")

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			list = append(list, Data{fname: checkChar20(parts[0]), lname: checkChar20(parts[1])})
		}
	}

	for _, d := range list {
		fmt.Println("Firstname is " + d.fname + " Lastname is " + d.lname)
	}

}

func checkChar20(str string) string {
	if len(str) > 20 {
		return str[:20]
	}
	return str
}
