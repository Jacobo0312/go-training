package structs

import (
	"encoding/json"
	"fmt"
)

// Global Slice for Employees

var employees []Employee

// Employee struct
// ID, Position, Person.
// Tags: Use json for marshalling and unmarshalling
type Employee struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Person   Person `json:"person"`
}

//Methods

func (e *Employee) PrintEmployee() {
	fmt.Println("ID: ", e.ID)
	fmt.Println("Position: ", e.Position)
	e.Person.Print()
}

func (e *Employee) PrintEmployeeJson() {
	jsonEmployee, _ := json.Marshal(e)
	fmt.Println(string(jsonEmployee))
}
