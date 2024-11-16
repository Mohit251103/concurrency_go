package main

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

type philosophers struct {
	id int
}
