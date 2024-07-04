package main

import "fmt"

func main() {

	// Arrays - Vectors

	// Declaring an array of integers

	// Method 1
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50

	fmt.Println(arr1)

	// Method 2
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	// Method 3
	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr3)

	// Printing the length of the array

	fmt.Println("Length of arr1:", len(arr1))

	// Example with strings

	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	// Example Ellipsis

	// The ellipsis operator (...) is used to create an array with a length based on the number of elements in the initializer list.
	// Position 99 will be -1, and the rest will be 0
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	// Multidimensional arrays

	// Declaring a 2D array

	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println(matrix)

}
