package main

import (
	"fmt"
	"runtime" // import runtime package for implementing parallelism
	"sync"    // package used for synchronization
	"time"
)

func main() {
	fmt.Println("Concurrency in GO!")

	// Call runtime.GOMAXPROCS to up the number of virtual processes or threads available to our program!
	runtime.GOMAXPROCS(2)

	// What is Concurrency?
	// Creating multiple processes that execute independently

	// Concurrency (independently) vs Parallelism (simultaneously)
	// http://blog.golang.org/concurrency-is-not-parallelism
	// Concurrency is the composition of independently executing
	// processes, while parallelism is the simultaneous execution of
	// (possibliy related) computations. Concurrency is about dealing with
	// lots of things at once. Parallelism is about doing lots of things at one - Rob Pike

	// Go's Concurrency Model
	// goroutines - all action happens here , this is a light weight abstraction on top of threads.
	// These are scheduled by the Go Runtime. (something like TPL on top of threading in .NET)
	// Lightweight - few KBs
	// Go manages goroutines (simpler for programmers)
	// Less switching (better performance) Runtime handles when a thread is blocked on I/O
	// Faster start-up times.
	// Safe communication between each other goroutines

	// Go's Concurrency Model is Actor Model - Communicating Sequential Processes - CSP
	// Actors safely pass message between each other via channels. One goroutine puts it in the channel and
	// another goroutine grabs it from the channel. Some kind of a pipe.

	// Write some Concurrent Code
	// Below are examples of anonymous functions
	func() {
		time.Sleep(5 * time.Second) // The entire program is blocked and we don't do anything for 5 seconds.
		fmt.Println("Hello")
	}() // Self executing function!

	func() {
		fmt.Println("Concurrency")
	}() // Self executing function!

	var waitGroup sync.WaitGroup
	// Add the number to the counter. If the counter becomes zero, all goroutines blocked on Wait are released.
	waitGroup.Add(2)

	// Adding the 'go' keyword in front of the routine makes them go routines.
	go func() {
		defer waitGroup.Done() // When done is called it internally decrements the waitgroup - wg.Add(-1)! :)

		time.Sleep(5 * time.Second) // Just this goroutine is blocked!
		fmt.Println("Concurrent Hello from Go")
	}() // Self executing function!

	go func() {
		defer waitGroup.Done()

		fmt.Println("Concurrency!!!")
	}() // Self executing function!

	// We need to wait for the goroutines to complete! Otherwise the program simply exits!
	waitGroup.Wait()

	// Channels can be either buffered or unbuffered!
	// Unbuffered channel is the one that we create using the make and chan keyword, but we don't specify a buffer size.
	// myNewChannel := make(chan int)
	// This type of an unbuffered channel cannot hold any item on its own and it is just like a pipe for data to flow through!
	// One goroutine putting data into the channel should wait for another goroutine to come and collect the data.
	// The first goroutine locks until there is another goroutine to pick the data, once picked the original goroutine moves on!

	// Buffered channel on the other hand, is created by providing a size parameter.
	// myNewBufferedChannel := make(chan int, 5)
	// Size denotes the number of data elements the channel can hold.

	// The goroutine can put the data in the channel and proceed further. However, if the channel is full,
	// then it is blocked until a buffer is available.

	// Similarly you are trying to get data off a channel and if there is no data then you are locked!
}
