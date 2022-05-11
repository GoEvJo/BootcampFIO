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
	minor int
	for _, num:=range numbers{
		if numbers < minor {
			minor=numbers
		}
	}
	return numbers
}

func main() {
	number := product(5 * 5)
	fmt.Println("El producto es: ", number)
	fmt.Println(minorOf(5,89,24,64,22,1))
}
