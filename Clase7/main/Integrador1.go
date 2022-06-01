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
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxID = 100
const fileName = "customers.txt"

var noFile = errors.New("el archivo indicado no fue encontrado o está dañado")

var flagError bool = false

type client struct {
	ID           int
	DNI          int
	nameLastname string
	telNumber    int
	address      string
}

func (c *client) showClient() {
	fmt.Println("\nA continuación, los datos del cliente...")
	fmt.Println("Nº de cuenta: ", c.ID)
	fmt.Println("Nombre: ", c.nameLastname)
	fmt.Println("DNI: ", c.DNI)
	fmt.Println("Tel.: ", c.telNumber)
	fmt.Println("Dirección: ", c.address, "\n")
}

func doesTheFileExist(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return noFile
	}
	return nil
}

func idGenerator(idGenerated *int, fileName string) (int, error) {
	if err := doesTheFileExist(fileName); err != nil {
		return 0, err
	}
	for *idGenerated <= maxID {
		checkField, checkError := checkExisting(1, *idGenerated, fileName)
		if checkError != nil {
			return 0, checkError
		}
		if checkField != false {
			break
		}
		*idGenerated++
	}
	if *idGenerated > maxID {
		return 0, fmt.Errorf("el ID generado supera el valor permitido permitido")
	}
	return *idGenerated, nil
}

func checkExisting(field, data int, file2Open string) (checkField bool, checkError error) {
	if errors.Is(noFile, doesTheFileExist(file2Open)) {
		return false, noFile
	}

	newDataString := strconv.Itoa(data)
	newDataString2 := "\r\n" + newDataString
	toRead, err := os.ReadFile(file2Open)
	if err != nil {
		return false, err
	}
	input := string(toRead)
	control := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	onSemicolon := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ';' {
				control++
				if control == 6 {
					control = 1
				}
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onSemicolon)
	// Scan.
	for scanner.Scan() {
		if control == field {
			myText := scanner.Text()
			if myText == newDataString || myText == newDataString2 {
				checkField = false
				checkError = nil
				return checkField, checkError
			} else {
				checkField = true
				checkError = nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return checkField, checkError
}

func checkingFields(id, dni, tel int, name, address string) error {
	if id == 0 || dni == 0 || tel == 0 || name == "" || address == "" {
		return fmt.Errorf("uno de los campos se encuentra en blanco")
	} else {
		return nil
	}
}

func sayGoodbye() {
	if flagError == true {
		fmt.Println("Se detectaron varios errores en tiempos de ejecución.")
	} else {
		fmt.Println("Fin de la ejecución de manera correcta.")
	}
}

func main() {
	defer sayGoodbye()

	if errors.Is(noFile, doesTheFileExist(fileName)) {
		flagError = true
		panic(noFile)
	}

	idCount := 1
	newID, idErr := idGenerator(&idCount, fileName)
	if idErr != nil {
		flagError = true
		panic(idErr)
	}
	newDNI := 2711848
	checkingDNI, errDNI := checkExisting(2, newDNI, fileName)
	if errDNI != nil {
		fmt.Println(errDNI)
	}
	if checkingDNI != true {
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
}
