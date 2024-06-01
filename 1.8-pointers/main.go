package main

import (
	"fmt"
)

func main() {

	// Pointers

	fmt.Println("Pointers")

	var a int = 42
	var b *int = &a

	fmt.Println("Memory Address of a:", &a)

	fmt.Println(a, *b)

	a = 27
	fmt.Println(a, *b)

	*b = 14
	fmt.Println(a, *b)
}
