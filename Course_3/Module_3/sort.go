package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// แบ่งอาร์เรย์ออกเป็น 4 ส่วน
	numParts := 4
	partSize := n / numParts
	remainder := n % numParts

	parts := make([][]int, numParts)
	start := 0

	for i := 0; i < numParts; i++ {
		end := start + partSize
		if i < remainder {
			end++
		}
		parts[i] = arr[start:end]
		start = end
	}
	var wg sync.WaitGroup
	for i := 0; i < numParts; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d sorting: %v\n", i+1, parts[i])
			sort.Ints(parts[i])
			fmt.Printf("Goroutine %d sorted: %v\n", i+1, parts[i])
		}(i)
	}

	wg.Wait()

	sortedArray := mergeSortedParts(parts)
	fmt.Println("Final sorted array:", sortedArray)
}

func mergeSortedParts(parts [][]int) []int {
	result := make([]int, 0)
	indices := make([]int, len(parts))

	for {
		minVal := int(^uint(0) >> 1)
		minIndex := -1

		for i, part := range parts {
			if indices[i] < len(part) && part[indices[i]] < minVal {
				minVal = part[indices[i]]
				minIndex = i
			}
		}

		if minIndex == -1 {
			break
		}

		result = append(result, minVal)
		indices[minIndex]++
	}

	return result
}
