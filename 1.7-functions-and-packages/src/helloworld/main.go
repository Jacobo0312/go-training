package main

import (
	"fmt"
	"github.com/jacobo0312/calculator"
)

func main() {
	a := 10
	b := 5
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	sum := sum(a, b)
	fmt.Println("Sum:", sum)
	diff := subtract(a, b)
	fmt.Println("Difference:", diff)
	sumResult, mulResult := calc(a, b)
	fmt.Println("Sum and Product:", sumResult, mulResult)

	//Package calculator

	total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)
}

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func subtract(number1 int, number2 int) (result int) {
	result = number1 - number2
	return
}

func calc(number1 int, number2 int) (div int, mul int) {
	int1 := number1
	int2 := number2
	div = int1 / int2
	mul = int1 * int2
	return
}
