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
When main() returns the program exits and takes the boring function down with it.
So this go routine never gets to do anything.
*/
func main() {
	go boring("jmc")
}
