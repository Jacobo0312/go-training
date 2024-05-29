package main

import "fmt"

func main() {

	// Declare a single variable
	//var firstName string

	// Declare multiple variables
	//var firstName, lastName string

	// Declare multiple variables with different types
	/*
		var firstName, lastName string
		var age int
	*/

	// Declare multiple variables with different types using shorthand
	/*
		var (
			firstName, lastName string
			age                 int
		)
	*/

	// Declare multiple variables with different types and initialize
	/*
		var (
			firstName string = "Jacobo"
			lastName  string = "Garcia"
			age       int    = 21
		)
	*/

	// Declare multiple variables with different types and initialize using shorthand

	firstName, lastName, age := "Jacobo", "Garcia", 21

	firstName = "Juan"

	// Constants

	const HTTPStatusOK = 200

	// Constants with multiple values

	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	fmt.Println(firstName, lastName, age)

	fmt.Println(HTTPStatusOK)

	fmt.Println(StatusOK, StatusConnectionReset, StatusOtherError)

}
