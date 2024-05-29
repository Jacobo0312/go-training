package main

import (
	"fmt"
	"math"
)

func main() {
	/* Basic Operations:
	   1. Arithmetic Operations
	   2. Relational Operations
	   3. Logical Operations
	   4. Bitwise Operations
	   5. Assignment Operations
	   6. Misc Operations
	*/

	// Arithmetic Operations
	var a int = 10
	var b int = 20

	sum := a + b
	diff := a - b
	product := a * b
	quotient := b / a
	remainder := b % a
	unaryPlus := +a
	unaryMinus := -b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	fmt.Println("Unary +:", unaryPlus)
	fmt.Println("Unary -:", unaryMinus)

	// Relational Operations
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b
	isEqual := a == b
	isNotEqual := a != b

	fmt.Println("Greater than:", isGreater)
	fmt.Println("Less than:", isLess)
	fmt.Println("Greater or equal:", isGreaterOrEqual)
	fmt.Println("Less or equal:", isLessOrEqual)
	fmt.Println("Equal to:", isEqual)
	fmt.Println("Not equal to:", isNotEqual)

	// Logical Operations
	logicalAND := true && false
	logicalOR := true || false
	logicalNOT := !true

	fmt.Println("Logical AND:", logicalAND)
	fmt.Println("Logical OR:", logicalOR)
	fmt.Println("Logical NOT:", logicalNOT)

	// Bitwise Operations
	x := 0b00101010 // Binary for 42
	y := 0b00001101 // Binary for 13

	bitwiseAND := x & y
	bitwiseOR := x | y
	bitwiseXOR := x ^ y
	bitwiseComplement := ^x

	fmt.Printf("Bitwise AND: %08b\n", bitwiseAND)
	fmt.Printf("Bitwise OR: %08b\n", bitwiseOR)
	fmt.Printf("Bitwise XOR: %08b\n", bitwiseXOR)
	fmt.Printf("Bitwise Complement: %08b\n", bitwiseComplement)

	// Assignment Operations
	var c int = 5
	c += 3 // c = c + 3
	c -= 2 // c = c - 2
	c *= 4 // c = c * 4
	c /= 2 // c = c / 2
	c %= 3 // c = c % 3

	fmt.Println("Final value of c:", c)

	// Misc Operations
	d := 10
	e := 3

	fmt.Println("10 raised to 3:", math.Pow(float64(d), float64(e)))
	fmt.Println("Absolute value of -5:", math.Abs(-5))
	fmt.Println("Maximum of 10 and 20:", math.Max(float64(d), 20))
	fmt.Println("Minimum of 10 and 20:", math.Min(float64(d), 20))
}
