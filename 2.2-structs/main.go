package main

import "fmt"

// Structs

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {

	// var Jacobo Employee

	// Jacobo.ID = 1
	// Jacobo.FirstName = "Jacobo"
	// Jacobo.LastName = "Garcia"
	// Jacobo.Address = "Calle 1, Ciudad 1"

	Jacobo := Employee{1, "Jacobo", "Garcia", "Calle 1, Ciudad 1"}

	// Printing the struct
	fmt.Println(Jacobo)

	// Copy reference
	JacoboCopy := &Jacobo

	JacoboCopy.LastName = "Aristizabal"

	fmt.Println(JacoboCopy)

}
