package main

import (
	"fmt"
	"time"
)

// goroutines are functions or subroutines which run on completely different thread then the main thread in a go prog
// each goroutine is assigned to one OS thread. Goroutines are lightweight and less resource intensive.

// channels in go are used to establish a two way communication medium between two goroutines
// each channel can send only one type of data
// by defauly synchronized, so no need for synchronizing explicitly or it can't be done ig :p
// chan keyword
// <- for either sending or receiving a data to goroutines

// receiver
func sync_print(ch chan int) {
	for val := range ch {
		fmt.Println("Running thread", val)
		time.Sleep(2 * time.Second)
		fmt.Println("Ended thread", val)
	}
}

func channel_main() {

	ch := make(chan int) // or var chan int   This is an unbuffered channel so cannot hold values if the receiver is not up
	// ch := make(chan int, 10)  ---- this one is a buffered channel with buffer size of 10
	go sync_print(ch)

	for i := 0; i < 10; i++ { // sender
		ch <- i
	}

	close(ch) // not necessary in this case but is a good practice to adopt

}
