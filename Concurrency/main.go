package main

import (
	"fmt"
	"sync"
)

func main() {
	// HellloRoutine()
	// CalculateParallel()
	TaskQueue()
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


//if you need to send something into channel use <-
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

	totalsum := 0
	for  v := range ch {
		totalsum += v		
	}
	fmt.Println(totalsum)
}

//when use need to use value in channel do not use <-  
func TaskWorker(order int, task chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker " + fmt.Sprint(order+1) + " processing " + <-task)
}

func TaskQueue()  {
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	
	var wg sync.WaitGroup
	ch := make(chan string,len(tasks))

	for i := range tasks {
		wg.Add(1)
		go TaskWorker(i,ch,&wg)
	}

	for _, v := range tasks {
		ch <- v	
	}

	wg.Wait()

	fmt.Println("All worker done !")

	
}