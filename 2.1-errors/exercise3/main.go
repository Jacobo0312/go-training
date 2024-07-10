package main

import (
	"errors"
	"fmt"
)

// Errors
var (
	ErrIDRequired              = errors.New("Error: ID is required")
	ErrNameRequired            = errors.New("Error: Name is required")
	ErrPhoneNumberRequired     = errors.New("Error: PhoneNumber is required")
	ErrHomeRequired            = errors.New("Error: Home is required")
	ErrFileRequired            = errors.New("Error: File is required")
	CustomerAlreadyExistsError = errors.New("Error: customer already exists")
)

// Global Slice for Customers
var Customers []Customer

type Customer struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Home        string
}

// Validate the customer
func (c *Customer) Validate() error {
	if c.ID == 0 {
		return ErrIDRequired
	}
	if c.Name == "" {
		return ErrNameRequired
	}
	if c.PhoneNumber == "" {
		return ErrPhoneNumberRequired
	}
	if c.Home == "" {
		return ErrHomeRequired
	}
	if c.File == "" {
		return ErrFileRequired
	}
	return nil
}

func (c *Customer) Save() error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime", r)
		}
	}()

	// Validate the customer
	if err := c.Validate(); err != nil {
		return err
	}

	// Check if the customer already exists
	for _, customer := range Customers {
		if customer.ID == c.ID {
			panic(CustomerAlreadyExistsError)
		}
	}

	// Save the customer
	Customers = append(Customers, *c)
	return nil
}

func main() {

	Customers1 := Customer{
		File:        "customer1.txt",
		Name:        "Jacobo Garcia",
		ID:          1,
		PhoneNumber: "1234567890",
		Home:        "123 Main St",
	}

	Customers2 := Customer{
		File:        "customer2.txt",
		Name:        "Jacobo Garcia",
		PhoneNumber: "1234567890",
		Home:        "123 Main St",
	}

	err := Customers1.Save()
	if err != nil {
		fmt.Println(err)
	}

	err = Customers1.Save()
	if err != nil {
		fmt.Println(err)
	}

	err = Customers2.Save()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Customers)

	fmt.Println("End of execution")

}
