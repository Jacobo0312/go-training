package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Launch a goroutine that simulates a task that could take longer than expected
	go func(ctx context.Context) {
		// Simulate a task that could take longer than expected
		select {
		case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
			fmt.Println("Task completed")
		case <-ctx.Done():
			fmt.Println("Cancelled task", ctx.Err())
		}
	}(ctx)

	// Wait for a while to allow the goroutine to run
	time.Sleep(500 * time.Millisecond)

	// Cancel the context to signal the goroutine to stop
	cancel()

	// Wait for a while to allow the goroutine to stop
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Finishing main")
}
