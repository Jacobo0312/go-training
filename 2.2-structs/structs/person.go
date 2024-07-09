package structs

import (
	"fmt"
	"encoding/json"
)

// Global Slice for Persons

// var persons []Person

// Person struct
//	ID, Name, DateOfBirth.
// Tags: Use json for marshalling and unmarshalling
type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

//Methods

// Print in string format
func (p *Person) Print() {
	fmt.Println("ID: ", p.ID)
	fmt.Println("Name: ", p.Name)
	fmt.Println("Date of Birth: ", p.DateOfBirth)
}

// Print in json format
func (p *Person) PrintJson() {
	jsonPerson, _ := json.Marshal(p)
	fmt.Println(string(jsonPerson))
}
