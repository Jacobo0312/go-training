package main

import (
	"fmt"
	"regexp"
)

func giveMeNumber() int {
	return 30
}

func main() {

	// If Statements
	fmt.Println("If Statements")

	if num := giveMeNumber(); num > 50 {
		fmt.Println(num, "is greater than 50")
	} else if num == 50 {
		fmt.Println(num, "is equal to 50")
	} else {
		fmt.Println(num, "is less than 50")
	}

	// Switch Statements

	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	fmt.Println("Switch Statements")
	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}



}
