package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

/*
We can hang around a little and on the way show that both main and the launched goroutine are running.
*/
func main() {
	go boring("jmc")
	fmt.Println("I'm listening.")
	time.Sleep(5 * time.Second)
	fmt.Println("You're boring; I'm leaving.")

}
