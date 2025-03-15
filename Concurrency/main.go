package main

import (
	"fmt"
	"sync"
)

func main() {
	HellloRoutine()
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
