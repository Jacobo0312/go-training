package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var (
	ErrSalaryLessThan100000 = errors.New("Error: salary is less than 100000")
	ErrSalaryBelowTaxable   = errors.New("Error: the salary entered does not reach the taxable minimum")
)

// Errors
func main() {
	// Error handling
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	error_1 := errors.New("Error 1")
	error_2 := fmt.Errorf("Second error: %w", error_1)

	//Unwrap
	err_1 := errors.Unwrap(error_2)
	//Is
	if errors.Is(err_1, error_1) {
		fmt.Println("Error 1 found")
	}

	//As
	_, errOpenFile := os.Open("not_exist.txt")

	var pathError *fs.PathError
	if errors.As(errOpenFile, &pathError) {
		fmt.Println("File Path:", pathError.Path)
		fmt.Println("Op:", pathError.Op)
		fmt.Println("Err:", pathError.Err)
	}

	fmt.Println(">>> Start")
	// Panic
	causePanic()
	fmt.Println(">>> End")

	salary := 100000

	errSalary := checkSalary(salary)
	if errSalary != nil {
		fmt.Println(errSalary)
	}

	fmt.Println(errors.Is(errSalary, ErrSalaryLessThan100000))

	hours := 90.0
	value := 2000.0

	salaryCalculated, errCalculateSalary := calculateSalary(hours, value)

	if errCalculateSalary != nil {
		fmt.Println(errCalculateSalary)
	} else {
		fmt.Println("Salary calculated:", salaryCalculated)
	}

}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering", r)
		}
	}()
	// Panic
	panic("Panic!")
}

func checkSalary(salary int) error {
	if salary <= 100000 {
		return ErrSalaryLessThan100000
	} else if salary < 150000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}
	fmt.Println("Must pay tax")
	return nil
}

func calculateSalary(hours, value float64) (float64, error) {
	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	salary := hours * value
	if salary >= 150000 {
		return salary * 0.9, nil
	}
	return salary, nil
}
