package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers []int

	fmt.Print("Enter input 10 Numbers > ")
	input, _ := reader.ReadString('\n')

	nums := strings.Fields(input)

	for i, n := range nums {
		if i > 9 {
			break
		}
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)

	fmt.Println("Sorted integers:", numbers)
}

