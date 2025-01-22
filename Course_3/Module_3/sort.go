package main

import "fmt"

func bubbleSort(arr []int32) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var input = []int32{1, 5, 98, 87, 4, 8, 9, 10, 25, 45, 78, 13, 45, 95, 62, 47}
	n := len(input)
	numParts := 4
	parts := n / numParts
	var output = [][]int32{}
	for i := 0; i < numParts; i++ {
		start := i * parts
		end := parts + start
		if i == n-1 {
			end = n
		}
		output = append(output, input[start:end])
		fmt.Printf("\nPart: %d > ",i+1)
		fmt.Print(output[i])
		fmt.Print(" > ")
		go bubbleSort(output[i])
		fmt.Print(output[i])
	}

}