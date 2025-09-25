// Run code by using "go run firstpart.go"
package main

import (
	"fmt"
	"sync"
	"time"
)

// producer generates numbers and sends them to a channel.
// It uses a WaitGroup to signal when it's finished.
func producer(numbers chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {
		// Simulate some work
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Producer: %d\n", i)
		numbers <- i // Send the number to the channel
	}
}

// consumer receives numbers from a channel and prints them.
// It uses a WaitGroup to signal when it's finished.
func consumer(numbers <-chan int, wg *sync.WaitGroup) {
	
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Receive the number from the channel. This will block until a number is available.
		number := <-numbers
		fmt.Printf("Consumer: %d\n", number)
		// Simulate some work
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	start := time.Now() // Start timing

	// Create a channel to pass integers. The channel is unbuffered,
	// meaning a sender will block until a receiver is ready.
	numbers := make(chan int)

	// A WaitGroup is used to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Add two goroutines to the WaitGroup.
	wg.Add(2)

	// Start the producer and consumer goroutines.
	go producer(numbers, &wg)
	go consumer(numbers, &wg)

	// Wait for all goroutines in the WaitGroup to complete.
	wg.Wait()

	fmt.Println("\nMain goroutine finished. All processes have terminated.")
	fmt.Printf("Execution time: %s\n", time.Since(start)) // Print elapsed time
}
