package main

/*
Generator: function that returns a channel
Channels are first-class values, just like strings or integers.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	nextFib := fibonacci()
	for i := 0; i < 1000000; i++ {
		fmt.Println(i, "	", <-nextFib, "	", time.Since(start))
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func fibonacci() <-chan int { // Returns receive-only channel of strings.
	c := make(chan int)
	go func() {
		//for previous, current := 0, 1; ; previous, current = previous+current, previous {
		for previous, current := 0, 1; ; previous, current = current, previous+current {
			c <- previous
		}
	}()
	return c // Return the channel to the caller.
}
