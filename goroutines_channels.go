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

func sync_print(ch chan int) {
	for val := range ch {
		fmt.Println("Running thread", val)
		time.Sleep(2 * time.Second)
		fmt.Println("Ended thread", val)
	}
}

func channel_main() {

	ch := make(chan int) // or var chan int
	go sync_print(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}

}
