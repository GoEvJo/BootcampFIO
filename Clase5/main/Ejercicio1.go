/*
Punteros, funciones y structs
A. Crear una función simple que permita multiplicar dos números y
devolver el resultado.
B. Crear una función que reciba un número variable de enteros como
argumento y devuelva el menor de ellos.
C. Crear una estructura Animal, y varias estructuras que tomen sus
propiedades y representen diferentes tipos de animales. Crear
además una función que permita instanciar un Animal y devuelva un
puntero al mismo.
D. Crear un a función que permita obtener de forma recursiva el
factorial de un número.
*/

package main

import "fmt"

func product(a, b int) int {
	result := a * b
	return result
}
func minorOf(numbers ...int) int {
	minor := 0
	for i, num := range numbers {
		if i == 0 {
			minor = num
		}
		if num < minor {
			minor = num
		}
	}
	return minor
}

type animal struct {
	kind            string
	paw, eyes, wing int
}
type pet struct {
	Name string
	animal
	Race, Color string
}

func /*(r pet)*/ load() *pet {
	NewPet := pet{
		Name: "Lolo",
		animal: animal{kind: "Dog",
			paw:  4,
			eyes: 2,
			wing: 0,
		},
		Race:  "Street",
		Color: "Black",
	}
	PointingNewPet := &NewPet
	return PointingNewPet
}

func factorial(a int) int64 {
	if a == 0 {
		return 1
	}
	return int64(a) * factorial(a-1)
}

func main() {
	number := product(5, 5)
	fmt.Println("El producto es: ", number)
	fmt.Println("El menor encontrado es: ", minorOf(5, 89, 24, 64, 22, 15))

	fmt.Println("La nueva mascota es: ", load())
}
