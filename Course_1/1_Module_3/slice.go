package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var input string
	s := make([]int, 3)
	i := 0

	for true {

		fmt.Println("slice updated: ", s)
		fmt.Print("Enter number: ")
		fmt.Scan(&input)

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "x" {
			return
		}

		if num, err := strconv.Atoi(input); err == nil {
			sort.Sort(sort.IntSlice(s))
			if i < 3 {
				s[i] = num
				i++
				continue
			}
			s = append(s, num)

		}
		i++

	}

}
