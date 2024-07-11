package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	FileNotFoundError = errors.New("The indicated file was not found or is damaged")
)

func main() {
	const path = "customers.txt"

	err := readFile(path)

	if err != nil {
		fmt.Println(err)
	}

}

func readFile(path string) error {
	file, err :=
		os.Open(path)
	if err != nil {
		return FileNotFoundError
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering", r)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
