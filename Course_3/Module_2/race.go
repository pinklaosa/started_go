package main

import (
	"fmt"
	"sync"
)

var count int64

func main() {
	
	var wg sync.WaitGroup

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 500; i++ {
			count++
		}

	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 500; i++ {
			count--
		}
	}(&wg)
	
	wg.Wait()
	fmt.Print(count)

}
