package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func showRandNum(n int) {
	fmt.Println("please wait for a while...")

	// sleep random time
	r := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(r.Intn(3000)))

	// get random number in [0, n)
	num := rand.Intn(n)
	fmt.Printf("random number is %d\n", num)
}

func goRoutineWaitGroup() {
	fmt.Println("run goRoutineWaitGroup()")

	// go showRandNum(10) // main goroutin will finish without waiting

	var wg sync.WaitGroup
	wg.Add(1) // increment counter by 1
	go func() {
		defer wg.Done() // decrement counter by one when func() return
		showRandNum(10)
	}()
	wg.Wait() // wait for wg until the counter is 0
}
