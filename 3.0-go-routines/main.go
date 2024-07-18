package main

import (
	"fmt"
	"sync"
	"time"
)

func makePasta(wg *sync.WaitGroup) {
	// Decrement the count when the goroutine completes
	defer wg.Done()
	// Display the pasta making process
	fmt.Println("Boiling water")
	fmt.Println("Cooking pasta")
}

func makeSauce(wg *sync.WaitGroup) {

	// Decrement the count when the goroutine completes
	defer wg.Done()
	fmt.Println("Heating olive oil")
	fmt.Println("Browning the garlic")
	fmt.Println("Browning the onions")
	fmt.Println("Cooking the meatballs")
	fmt.Println("Cooking the diced tomatoes")
}

//Create WaitGroup

var wg sync.WaitGroup

func main() {

	start := time.Now()
	// Add a count of two, one for each goroutine
	wg.Add(2)
	// Launch two goroutines
	go makePasta(&wg)
	go makeSauce(&wg)
	// Wait for the goroutines to finish
	wg.Wait()
	// Serve the pasta and sauce
	fmt.Println("Dinner is served!")

	duration := time.Since(start)

	fmt.Println(duration)

}
