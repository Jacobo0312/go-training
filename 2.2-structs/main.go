package main

import (
	"fmt"

	"github.com/Jacobo0312/go-training/2.2-structs/structs"
)

// Structs

// type Employee struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Address   string
// }

func main() {

	// var Jacobo Employee

	// Jacobo.ID = 1
	// Jacobo.FirstName = "Jacobo"
	// Jacobo.LastName = "Garcia"
	// Jacobo.Address = "Calle 1, Ciudad 1"

	// Jacobo := Employee{1, "Jacobo", "Garcia", "Calle 1, Ciudad 1"}

	// Printing the struct
	// fmt.Println(Jacobo)

	// // Copy reference
	// JacoboCopy := &Jacobo

	// JacoboCopy.LastName = "Aristizabal"

	// fmt.Println(JacoboCopy)

	// Creating a new product
	Product1 := structs.Product{
		ID:          1,
		Name:        "Product 1",
		Price:       100,
		Description: "Description 1",
		Category:    "Category 1",
	}

	Product2 := structs.Product{
		ID:          2,
		Name:        "Product 2",
		Price:       200,
		Description: "Description 2",
		Category:    "Category 2",
	}

	Product3 := structs.Product{
		ID:          3,
		Name:        "Product 3",
		Price:       300,
		Description: "Description 3",
		Category:    "Category 3",
	}

	// Saving the product
	Product1.Save()
	Product2.Save()
	Product3.Save()

	// Getting all the products
	Products := Product1.GetAll()

	// Printing the products

	fmt.Println(Products)

	// Finding a product by ID
	ProductFound := Product1.FindByID(2)

	fmt.Println(ProductFound)

	// Creating a new person
	Person1 := structs.Person{
		ID:          1,
		Name:        "Person 1",
		DateOfBirth: "01/01/2001",
	}

	// Creating a new employee

	Employee1 := structs.Employee{
		ID:       1,
		Position: "Position 1",
		Person:   Person1,
	}

	// Printing the person
	Employee1.Person.Print()

	// Printing the employee
	Employee1.PrintEmployee()

}
