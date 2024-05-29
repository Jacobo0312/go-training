package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		There are two types of data types in Go:
		1. Primitive Data Types
		2. Derived Data Types

		Primitive Data Types:
		1. Integer
		2. Float
		3. Complex
		4. Boolean
		5. String
		7. Array
		8. Structure

		Derived Data Types:
		1. Function
		2. Slice
		3. Interface
		4. Channel
		5. Map
		6. Pointer
	*/

	/*
		Integer Data Type:
		1. int8
		2. int16
		3. int32

		Note: int is an alias for int32 or int64 depending on the underlying platform.
	*/

	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807

	fmt.Println(integer8, integer16, integer32, integer64)

	/*
		Float Data Type:
		1. float32
		2. float64
	*/

	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807

	fmt.Println(float32, float64)
	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	/*
		Boolean Data Type:
		1. bool
	*/

	var featureFlag bool = true
	fmt.Println(featureFlag)

	/*
		String Data Type:
		1. string
			\n to new line
			\r to carriage return
			\t to tab
			\' to simple quote
			\" to double quote
			\\ to backslash
	*/

	var firstName string = "Jacobo"
	lastName := "Garcia"
	fmt.Println(firstName, lastName)

	fullName := "Jacobo Garcia \t(alias \"Jacob\")\n"
	fmt.Println(fullName)

	// Cast int16 to int32
	var a int16 = 127
	var b int32 = 32767
	fmt.Println(int32(a) + b)

	/*
		Uint Data Type:
		Integer data type with no sign.
		1. uint8
		2. uint16
		3. uint32
		4. uint64
	*/

	var uint8 uint8 = 255
	var uint16 uint16 = 65535
	var uint32 uint32 = 4294967295
	var uint64 uint64 = 18446744073709551615

	fmt.Println(uint8, uint16, uint32, uint64)

}
