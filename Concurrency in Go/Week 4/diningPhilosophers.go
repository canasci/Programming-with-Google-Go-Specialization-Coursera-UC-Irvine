package main

import (
	"fmt"
	"sync"
)

/*
[1] There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
[2] Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
[3] The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
[4] In order to eat, a philosopher must get permission from a host which executes in its own goroutine. The host allows no more than 2 philosophers to eat concurrently.
[5] Each philosopher is numbered, 1 through 5.
[6] When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
[7] When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

// Setup executes when Philosophers start dining
func Setup() {
	fmt.Printf("Dinner starts...\n\n")
}

// ChopStick struct acts a Mutex since access to them
type ChopStick struct{ sync.Mutex }

// Philosopher struct, has a unique number field, pointers to surrounding chopsticks assigned as two data fields
// Pointers to ChopStick object are given to alter the Mutex states (lock/ unlock)
type Philosopher struct {
	number          int
	eatCount        int
	leftCS, rightCS *ChopStick
}

var on sync.Once

var mut = &sync.Mutex{}

var totalCount int

// Eat method for Philosopher struct
func (p *Philosopher) Eat(wg *sync.WaitGroup, c chan *Philosopher) {
	on.Do(Setup)
	for i := 0; i < 3; i++ {
		if p.eatCount < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			p.eatCount++
			fmt.Printf("Starting to eat <%d>\n", p.number)
			fmt.Printf("Finishing eating <%d>\n", p.number)
			fmt.Printf("The total amount of times eaten by Philosopher <%d>: %d\n", p.number, p.eatCount)

			mut.Lock()
			totalCount++
			fmt.Printf("The total amount of times eaten by all Philosophers: %d\n\n", totalCount)
			mut.Unlock()

			p.leftCS.Unlock()
			p.rightCS.Unlock()

			c <- p
		}
	}
	if p.eatCount == 3 {
		wg.Done()
	}
}

// Host function
func Host(wg *sync.WaitGroup, c chan *Philosopher) {
	for {
		mut.Lock()
		if len(c) == 2 {
			<-c
			<-c
		} else if totalCount == 15 { // Each philosopher eats 3 times, there are 5 philosophers thus dinner ends when total count is 15
			fmt.Printf("Dinner ends...\n")
			break
		}
		mut.Unlock()
	}
	wg.Done()
}

func main() {
	c1 := make(chan *Philosopher, 2) // buffered channel; in order to contain 2 pointer to philosopher objects in transit
	var wg sync.WaitGroup            // Waitgroup object

	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, 0, chopSticks[i], chopSticks[(i+1)%5]}
	}

	wg.Add(1)
	go Host(&wg, c1)

	for j := 0; j < 5; j++ {
		wg.Add(1)
		go philosophers[j].Eat(&wg, c1)
	}

	wg.Wait()
}
