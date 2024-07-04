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
	numbersEllipsis := [...]int{99: -1}
	fmt.Println("First Position:", numbersEllipsis[0])
	fmt.Println("Last Position:", numbersEllipsis[99])
	fmt.Println("Length:", len(numbersEllipsis))

	// Multidimensional arrays

	// Declaring a 2D array

	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println(matrix)

	//Slices

	fmt.Println("Slices")

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	// Slicing

	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	// Append

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}

	// Remove

	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}

	// Copy
	letters = []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

	//Example with make

	// The make function is used to create slices. It takes three arguments: the type of the slice, the length, and the capacity.
	slice := make([]int, 3, 5)
	fmt.Println(slice, len(slice), cap(slice))

	//Map

	fmt.Println("Maps")

	// Declaring a map

	// Method 1

	var m map[string]int
	fmt.Println(m)

	// Method 2

	m = make(map[string]int)

	m["one"] = 1

	fmt.Println(m)

	// Method 3

	m = map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println(m)

	// Accessing a map

	fmt.Println("Accessing a map")

	fmt.Println(m["one"])

	// Deleting a key

	delete(m, "one")

	fmt.Println(m)

	// Checking if a key exists

	value, exists := m["one"]

	if exists {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key does not exist")
	}

	// Iterating over a map

	for key, value := range m {
		fmt.Println(key, value)
	}

}
