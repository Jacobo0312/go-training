package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
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
