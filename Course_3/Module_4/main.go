package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const maxMeals = 3

type Philosopher struct {
	id         int
	mealsEaten int
}

func main() {

	chopsticks := make([]sync.Mutex, numPhilosophers)
	host := make(chan struct{}, 2)

	var wg sync.WaitGroup

	philosophers := make([]Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{id: i + 1}
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go func(p *Philosopher, leftChopstick, rightChopstick *sync.Mutex) {
			defer wg.Done()

			for p.mealsEaten < maxMeals {
				host <- struct{}{}

				leftChopstick.Lock()
				rightChopstick.Lock()

				fmt.Printf("starting to eat %d\n", p.id)
				time.Sleep(time.Millisecond * 100)

				fmt.Printf("finishing eating %d\n", p.id)
				p.mealsEaten++

				rightChopstick.Unlock()
				leftChopstick.Unlock()
				<-host
			}
		}(&philosophers[i], &chopsticks[i], &chopsticks[(i+1)%numPhilosophers])
	}
	wg.Wait()
}
