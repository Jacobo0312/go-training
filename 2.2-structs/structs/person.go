package structs

import "fmt"

// Global Slice for Persons

var persons []Person

// Person struct
//  ID, Name, DateOfBirth.

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

//Methods

// Print in string format
func (p *Person) Print() {
	fmt.Println("ID: ", p.ID)
	fmt.Println("Name: ", p.Name)
	fmt.Println("Date of Birth: ", p.DateOfBirth)
}
