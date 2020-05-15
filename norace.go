package main

import (
	"fmt"
	"sync"
)

func div(v1 int, v2 int) int {
	return v1 / v2
}

func alterVar(p1 *int, wg *sync.WaitGroup) {
	*p1 = (*p1) * 2
	wg.Done()
}

func main() {
	a := 20
	b := 10
	var wg sync.WaitGroup // Waitgroup declared
	wg.Add(1)             // Goroutine added to the waitgroup in order to prevent data race by synchronizing the goroutines
	go alterVar(&a, &wg)  // Alters variable "a" to be two times its previous value
	wg.Wait()             // Waitgroup waits executing the main goroutine until alterVar goroutine executes correctly
	fmt.Println(div(a, b))
}
