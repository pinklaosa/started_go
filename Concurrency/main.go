package main

import (
	"fmt"
	"sync"
)

func main() {
	// HellloRoutine()
	// CalculateParallel()
	// TaskQueue()
	// PipelineProcessing()
	MathPowParallel()
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

// if you need to send something into channel use <-
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
	go SumSlices(numbers1, &wg, ch)
	go SumSlices(numbers2, &wg, ch)

	wg.Wait()
	close(ch)

	totalsum := 0
	for v := range ch {
		totalsum += v
	}
	fmt.Println(totalsum)
}

// when use need to use value in channel do not use <-
func TaskWorker(order int, task chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker " + fmt.Sprint(order+1) + " processing " + <-task)
}

func TaskQueue() {
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	var wg sync.WaitGroup
	ch := make(chan string, len(tasks))

	for i := range tasks {
		wg.Add(1)
		go TaskWorker(i, ch, &wg)
	}

	for _, v := range tasks {
		ch <- v
	}

	wg.Wait()

	fmt.Println("All worker done !")
}

func PipelineWorker(tasks chan int, totalSum chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for task := range tasks {
		sum += task
	}
	totalSum <- sum
}

func PipelineProcessing() {
	max := 10
	ch := make(chan int, max)
	for v := range 10 {
		ch <- v
	}
	close(ch)

	var wg sync.WaitGroup
	w := 3
	totalSum := make(chan int, w)

	for range w {
		wg.Add(1)
		go PipelineWorker(ch, totalSum, &wg)
	}

	wg.Wait()
	close(totalSum)

	finalSum := 0
	for partial := range totalSum {
		finalSum += partial
	}

	fmt.Println("All worker done: " + fmt.Sprint(finalSum))
}

func MathPowWorker(pows chan int, worker chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for v := range pows {
		sum += v
	}
	worker <- sum
}

func MathPowCal(pow chan <- int,num int,wg *sync.WaitGroup){
	defer wg.Done()
	pow <- num * num
}

func MathPowParallel() {
	n := 11
	w := 3
	nums := make(chan int, n)
	pows := make(chan int,n)
	worker := make(chan int, w)

	var wg sync.WaitGroup

	for v := range n {
		nums <- v
	}
	close(nums)

	for v := range nums {
		wg.Add(1)
		go MathPowCal(pows,v,&wg)
	}

	wg.Wait()
	close(pows)

	for range w {
		wg.Add(1)
		go MathPowWorker(pows, worker, &wg)
	}

	wg.Wait()
	close(worker)

	totalSum := 0
	for v := range worker {
		totalSum += v
	}

	fmt.Println(totalSum)
}
