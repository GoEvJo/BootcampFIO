/*
Clima
Una empresa de meteorología quiere tener una aplicación donde pueda
tener la temperatura y humedad y presión atmosférica de distintos lugares.
1. Declara 3 variables especificando el tipo de dato, como valor deben tener
la temperatura, humedad y presión de donde te encuentres.
2. Imprime los valores de las variables en consola.
3. ¿Qué tipo de dato le asignarías a las variables?

*/

package main

import "fmt"

func main() {
	var temperatura_C, humedad_porc, presion_hPa float64 = 15, 72, 1013.8

	fmt.Println("\nTemperatura : ", temperatura_C, "ºC")
	fmt.Println("Humedad : ", humedad_porc, "%")
	fmt.Println("Presión : ", presion_hPa, "hPa")
}
