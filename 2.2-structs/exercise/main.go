package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employees := []Employee{
		{
			Person: Person{
				LastName: "Garcia", FirstName: "Jacobo",
			},
		},
		{
			Person: Person{
				LastName: "Campaz", FirstName: "Camilo",
			},
		},
		{
			Person: Person{
				LastName: "Suarez", FirstName: "Gabriel",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)

}
