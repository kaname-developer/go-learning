package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandNum(n int, ch chan<- int) {
	fmt.Println("please wait for a while...")

	// sleep random time
	r := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(r.Intn(3000)))

	// get random number in [0, n)
	num := rand.Intn(n)

	ch <- num // send num to channel

}

func goRoutineChannel() {
	fmt.Println("run goRoutineChannel()")

	ch := make(chan int) // make channel which receives/sends int

	go getRandNum(10, ch)

	num := <-ch // wait for channel and receive num from it
	fmt.Printf("random number is %d\n", num)

	close(ch) // channel should be closed
}
