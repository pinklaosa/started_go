package main

import (
	"fmt"
	"sync"
)

func main() {
	// HellloRoutine()
	CalculateParallel()
}

func HellloRoutine() {
	n := 5
	var wg sync.WaitGroup
	for v := range n {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Hello from goroutine: " + fmt.Sprint(i))
		}(v, &wg)
	}
	wg.Wait()
	fmt.Println("All routine done !")
}

func SumSlices(numbers []int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	var sum int
	for _, v := range numbers {
		sum += v
	}
	ch <- sum
}

func CalculateParallel() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	partSize := len(numbers) / 2
	numbers1 := numbers[:partSize]
	numbers2 := numbers[partSize:]
	ch := make(chan int, 2)
	
	wg.Add(2)
	go SumSlices(numbers1,&wg,ch)
	go SumSlices(numbers2,&wg,ch)

	wg.Wait()
	close(ch)
	
	sum := 0
	for  v := range ch {
		sum += v		
	}
	fmt.Println(sum)
}
