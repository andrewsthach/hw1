package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 5 // keep same number as original

	start := time.Now()

	for i := 1; i <= N; i++ {
		// Producer work
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Producer: %d\n", i)

		// Consumer work
		fmt.Printf("Consumer: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nSequential (no goroutines) finished in %s\n", elapsed)
}
