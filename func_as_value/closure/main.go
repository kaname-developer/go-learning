package main

import (
	"fmt"
	"sync"
)

func countUpDownClosure() (f1 func(), f2 func(), f3 func() int) {
	count := 0

	f1 = func() {
		count++
	}

	f2 = func() {
		count--
	}

	f3 = func() int {
		return count
	}

	return
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	increaseCount, decreaseCount, getCount := countUpDownClosure()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()       // get lock
				increaseCount() // countup in each goroutine
				mu.Unlock()     // release lock
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50000; j++ {
				mu.Lock()
				decreaseCount() // countdown in each goroutine
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("count:", getCount()) // should be 150000
}
