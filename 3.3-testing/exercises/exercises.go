package exercises

import (
	"fmt"
)

/*
	Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

func CalculateTax(salary float64) float64 {
	if salary > 150000 {
		return salary * 0.27
	}
	if salary > 50000 {
		return salary * 0.17
	}
	return 0
}

/*
	Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

func CalculateAverage(grades ...int) float64 {
	sum := 0
	count := 0
	for _, grade := range grades {
		if grade < 0 {
			continue
		}
		sum += grade
		count++
	}
	return float64(sum) / float64(count)
}

/*

	Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	- Categoría C, su salario es de $1.000 por hora.
	- Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	- Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.

	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.

*/

func CalculateSalary(minutes int, category string) float64 {
	hours := float64(minutes) / 60
	switch category {
	case "C":
		return hours * 1000
	case "B":
		return hours * 1500 * 1.2
	case "A":
		return hours * 3000 * 1.5
	}
	return 0
}

/*

	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

	Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

*/

func Operation(op string) (func(...int) int, error) {
	switch op {
	case "minimum":
		return min, nil
	case "average":
		return ave, nil
	case "maximum":
		return max, nil
	}
	return nil, fmt.Errorf("Operation %s not defined", op)
}

func min(numbers ...int) int {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func ave(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

func max(numbers ...int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

/*
	Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

	Tienen los siguientes datos:

	Perro 10 kg de alimento.
	Gato 5 kg de alimento.
	Hamster 250 g de alimento.
	Tarántula 150 g de alimento.
	Se solicita:

	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/

var animals = map[string]func(amount int) float64{
	"dog":       dogFunc,
	"cat":       catFunc,
	"hamster":   hamsterFunc,
	"tarantula": tarantulaFunc,
}

func dogFunc(amount int) float64 {
	food := amount * 10000
	return float64(food)
}

func catFunc(amount int) float64 {
	food := amount * 5000
	return float64(food)
}

func hamsterFunc(amount int) float64 {
	food := amount * 250
	return float64(food)
}

func tarantulaFunc(amount int) float64 {
	food := amount * 150
	return float64(food)
}

func Animal(animalType string) (func(int) float64, string) {

	value, exists := animals[animalType]

	if !exists {
		return nil, "Animal not found"
	}
	return value, ""
}
