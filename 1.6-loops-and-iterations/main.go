package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Loop Statements

	// For Loop
	fmt.Println("For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While Loop
	fmt.Println("While Loop")
	var num int64
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	// Infinite Loop

	for {
		fmt.Println("Infinite Loop")
		break
	}

	// Continue Statement
	fmt.Println("Continue Statement")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Defer Statement

	fmt.Println("Defer Statement")

	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

}
