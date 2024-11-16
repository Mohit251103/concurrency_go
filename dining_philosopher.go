package main

import (
	"fmt"
	"sync"
)

// Problem statement :
// Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers.

// Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks.

// Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.

// Design a discipline of behaviour (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.

// Gist of the solution :
// We will try to run each philosopher's function as a separate goroutine
// This way we can achieve concurrency in the eating and thinking process

// For the shared resources i.e. forks in this case, we will try to handle their synchronized allocation and deallocation using channels
// So each fork will have a channel of its own ig which will be common across two philosophers

type DiningPhilosphers interface {
	eat()
	pickLeft() // fork
	pickRight()

	putLeft()
	putRight()
}

type philosophers struct {
	id        int
	forkLeft  chan struct{}
	forkRight chan struct{}
}

func (p *philosophers) pickLeft() {
	<-p.forkLeft
}

func (p *philosophers) pickRight() {
	<-p.forkRight
}

func (p *philosophers) putLeft() {
	p.forkLeft <- struct{}{}
}

func (p *philosophers) putRight() {
	p.forkRight <- struct{}{}
}

func (p *philosophers) eat(limit int, wg *sync.WaitGroup) {
	// pick the fork
	defer wg.Done()
	for i := 0; i < limit; i++ {
		fmt.Println("Philosopher", p.id, "is waiting for left and right fork")
		p.pickLeft()
		fmt.Println(p.id, "philospher got left fork")
		p.pickRight()
		fmt.Println(p.id, "philosopher got right fork")

		fmt.Println(p.id, "ate")

		// put the fork back
		p.putLeft()
		p.putRight()
	}
}

func start() {

	fork1 := make(chan struct{}, 1)
	fork2 := make(chan struct{}, 1)
	fork3 := make(chan struct{}, 1)
	fork4 := make(chan struct{}, 1)
	fork5 := make(chan struct{}, 1)

	// initializing the channels or basically making the forks available
	// struct{}{} here is a no data struct which does not occupy any memory and is solely used for signaling intent
	fork1 <- struct{}{}
	fork2 <- struct{}{}
	fork3 <- struct{}{}
	fork4 <- struct{}{}
	fork5 <- struct{}{}

	var table []philosophers

	p1 := philosophers{
		id:        1,
		forkLeft:  fork2,
		forkRight: fork1,
	}

	p2 := philosophers{
		id:        2,
		forkLeft:  fork3,
		forkRight: fork2,
	}

	p3 := philosophers{
		id:        3,
		forkLeft:  fork4,
		forkRight: fork3,
	}

	p4 := philosophers{
		id:        4,
		forkLeft:  fork5,
		forkRight: fork4,
	}

	p5 := philosophers{
		id:        5,
		forkLeft:  fork1,
		forkRight: fork5,
	}

	table = append(table, p1, p2, p3, p4, p5)

	var n int // How many times the philosophers could eat ??
	fmt.Println("Limit of process : ")
	fmt.Scanf("%d", &n)

	wg := sync.WaitGroup{}

	for _, philosopher := range table {
		wg.Add(1)
		go philosopher.eat(n, &wg)
	}

	wg.Wait()

}
