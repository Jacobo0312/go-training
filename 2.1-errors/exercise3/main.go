package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Errors
var (
	ErrIDRequired              = errors.New("Error: ID is required")
	ErrNameRequired            = errors.New("Error: Name is required")
	ErrPhoneNumberRequired     = errors.New("Error: PhoneNumber is required")
	ErrHomeRequired            = errors.New("Error: Home is required")
	ErrFileRequired            = errors.New("Error: File is required")
	CustomerAlreadyExistsError = errors.New("Error: customer already exists")
	FileNotFoundError          = errors.New("The indicated file was not found or is damaged")
	ErrFileIsEmpty             = errors.New("Error: File is empty")
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

	// Create a customer by file
	err := createUserByFile("customer1.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create a customer by file
	err = createUserByFile("customer1.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create a customer by file
	err = createUserByFile("customer2.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create a customer by file
	err = createUserByFile("customer3.txt")

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(Customers)

	fmt.Println("End of execution")

}

func createUserByFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return FileNotFoundError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return ErrFileIsEmpty
	}
	fileName := scanner.Text()

	if !scanner.Scan() {
		return ErrNameRequired
	}
	name := scanner.Text()

	if !scanner.Scan() {
		return ErrIDRequired
	}
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	if !scanner.Scan() {
		return ErrPhoneNumberRequired
	}
	phoneNumber := scanner.Text()

	if !scanner.Scan() {
		return ErrHomeRequired
	}
	home := scanner.Text()

	customer := Customer{
		File:        fileName,
		Name:        name,
		ID:          id,
		PhoneNumber: phoneNumber,
		Home:        home,
	}

	err = customer.Save()
	if err != nil {
		return fmt.Errorf("failed to save customer: %v", err)
	}

	return nil
}
