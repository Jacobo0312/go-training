package exercises_test

import (
	"testing"

	"github.com/Jacobo0312/go-training/3.3-testing/exercises"
	"github.com/stretchr/testify/require"
)

func TestCalculateTax_GreaterThan150000(t *testing.T) {

	//Arrange
	const (
		salary   = 152000
		expected = 41040.00
	)

	//Act
	result := exercises.CalculateTax(salary)

	//Assert
	require.Equal(t, expected, result)
}

func TestCalculateTax_GreaterThan50000(t *testing.T) {

	//Arrange
	const (
		salary   = 51000
		expected = 8670.00
	)

	//Act
	result := exercises.CalculateTax(salary)

	//Assert
	require.Equal(t, expected, result)
}


func TestCalculateAverage(t *testing.T) {

	//Arrange
	const (
		expected = 3.5
	)

	//Act
	result := exercises.CalculateAverage(4, 2, 3, 4, 5, 3)

	//Assert
	require.Equal(t, expected, result)
}


// Categoría C, su salario es de $1.000 por hora.
// Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
// Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.


func TestCalculateSalary_CategoryC(t *testing.T) {
	
	//Arrange
	const (
		minutes  = 60
		category = "C"
		expected = 1000.00
	)

	//Act
	result := exercises.CalculateSalary(minutes, category)

	//Assert
	require.Equal(t, expected, result)
}


func TestCalculateSalary_CategoryB(t *testing.T) {
	
	//Arrange
	const (
		minutes  = 60
		category = "B"
		expected = 1800.00
	)

	//Act
	result := exercises.CalculateSalary(minutes, category)

	//Assert
	require.Equal(t, expected, result)
}



func TestCalculateSalary_CategoryA(t *testing.T) {
	
	//Arrange
	const (
		minutes  = 60
		category = "A"
		expected = 4500.00
	)

	//Act
	result := exercises.CalculateSalary(minutes, category)

	//Assert
	require.Equal(t, expected, result)
}

func TestCalculateSalary_InvalidCategory(t *testing.T) {
	
	//Arrange
	const (
		minutes  = 60
		category = "D"
		expected = 0.00
	)

	//Act
	result := exercises.CalculateSalary(minutes, category)

	//Assert
	require.Equal(t, expected, result)
}






