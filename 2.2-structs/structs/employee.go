package structs

import "fmt"

// Global Slice for Employees

var employees []Employee

// Employee struct
// ID, Position, Person.
type Employee struct {
	ID       int
	Position string
	Person
}

//Methods

func (e *Employee) PrintEmployee() {

	fmt.Println("ID: ", e.ID)
	fmt.Println("Position: ", e.Position)
	e.Person.Print()

}
