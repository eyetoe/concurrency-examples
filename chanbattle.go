package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/util"
)

func KeyReader() {
	for {
		choice, _, _ := GetChar()
		// is this an int'able string? send on int channel
		if i, err := strconv.Atoi(choice); err == nil {
			IntInput <- i
		} else {
			// this should be non int'able alpha
			AlphaInput <- choice
		}
	}
}

func AlphaKeyPrinter() {
	for {
		fmt.Printf("Alpha: %q\n", <-AlphaInput)
	}

}

func IntKeyPrinter() {
	for {
		fmt.Printf("Int: %q\n", <-IntInput)
	}

}

var AlphaInput chan string
var IntInput chan int

func main() {
	AlphaInput = make(chan string)
	IntInput = make(chan int)

	go KeyReader()
	go AlphaKeyPrinter()
	go IntKeyPrinter()

	for {
		fmt.Printf("")
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
