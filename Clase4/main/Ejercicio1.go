/*
Arrays y slices
● Crear un string de 10 elementos del tipo entero.
● Recorrerlo y loguear los que son impares.
● Crear un slice que consista en los elementos [3,7] de
dicho string.
● Crear una función que reciba un slice de enteros y
devuelva su promedio.
*/

package main

import "fmt"

func average(data []int) float64 {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum = sum + data[i]
	}
	return float64(sum) / float64((len(data)))
}

func main() {
	/* Crear un Slice de 10 elementos enteros
	[!] Preguntar por una manera más fácil de inicializar un slice*/
	mySlice := make([]int, 10)
	mySlice[0] = 21
	mySlice[1] = 02
	mySlice[2] = 04
	mySlice[3] = 05
	mySlice[4] = 2022
	mySlice[5] = 4
	mySlice[6] = 1
	mySlice[7] = 1
	mySlice[8] = 18
	mySlice[9] = 8

	//Para loguear los impares:
	fmt.Println("Los impares son:")
	for i := 0; i < 10; i++ {
		if mySlice[i]%2 != 0 {
			fmt.Println(mySlice[i])
		}
	}

	//Crear otro Slice con los elementos 3 y 7 del anterior
	secondSlice := make([]int, 2)
	secondSlice[0] = mySlice[3]
	secondSlice[1] = mySlice[7]

	//Función que reciba un Slice y devuelva el promedio de sus valores
	fmt.Println("The average is: ", average(mySlice))

}
