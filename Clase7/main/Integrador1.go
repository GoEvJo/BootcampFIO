/*

Los datos requeridos para registrar a un cliente son:
- Legajo
- Nombre y Apellido
- DNI
- Número de teléfono
- Domicilio

*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const maxID = 100

var flagError bool = false

type client struct {
	ID           int
	DNI          int
	nameLastname string
	telNumber    int
	address      string
}

func (c *client) showClient() {
	fmt.Println("\nA continuaación, los datos del cliente...")
	fmt.Println("Nº de cuenta: ", c.ID)
	fmt.Println("Nombre: ", c.nameLastname)
	fmt.Println("DNI: ", c.DNI)
	fmt.Println("Tel.: ", c.telNumber)
	fmt.Println("Dirección: ", c.address, "\n")
}

func idGenerator() (int, error) {
	number := rand.Intn(maxID)
	if number <= 0 || number > maxID {
		return 0, fmt.Errorf("el ID se encuentra fuera del intervalo permitido")
	}
	return number, nil
}

func openFile(file2Open string) *os.File {
	myFile, err := os.Open(file2Open)
	if err != nil {
		flagError = true
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println("Archivo abierto exitosamente.")
	}
	return myFile
}

func closeFile(file2Close *os.File) {
	if file2Close.Close() != nil {
		fmt.Println("error al cerrar el archivo")
	} else {
		fmt.Println("Archivo cerrado de manera correcta.")
	}
}

func readFile(field, data int, file2Open string) (checkField bool) {
	file := openFile(file2Open)
	defer closeFile(file)
	/*
		Código para verificar existencia de datos
		...
		...
	*/
	return true
}

func checkingFields(id, dni, tel int, name, address string) error {
	if id == 0 || dni == 0 || tel == 0 || name == "" || address == "" {
		return fmt.Errorf("uno de los campos se encuentra en blanco")
	} else {
		return nil
	}
}

func sayGoodbye() {
	/*fmt.Println("Entrando a SayGoodBye con ", flagError)*/
	if flagError == true {
		fmt.Println("Se detectaron varios errores en tiempos de ejecución.")
	} else {
		fmt.Println("Fin de la ejecución de manera correcta.")
	}
}

func main() {
	defer sayGoodbye()
	rand.Seed(time.Now().UnixNano())

	var newID int
	var idErr error
	for {
		fmt.Println("Generando ID...")
		newID, idErr = idGenerator()
		if idErr != nil {
			flagError = true
			panic(idErr)
		}
		if readFile(1, newID, "customers.txt") == true {
			break
		}
	}
	newDNI := 2711848
	if readFile(2, newDNI, "customers.txt") != true {
		flagError = true
		panic("el cliente ya existe")
	}
	newNameLastname := "Boglioli Marce"
	newTelNumber := 27052022
	newAddress := "Calle Falsa 123"

	if errNil := checkingFields(newID, newDNI, newTelNumber, newNameLastname, newAddress); errNil != nil {
		fmt.Println(errNil)
	} else {
		newClient := client{
			ID:           newID,
			DNI:          newDNI,
			nameLastname: newNameLastname,
			telNumber:    newTelNumber,
			address:      newAddress,
		}
		fmt.Println("Nuevo cliente creado con éxito.")
		newClient.showClient()
	}

	/*flagError = true
	//fmt.Println("A ver... ", flagError)
	panic("panic at the disco!")*/
}
