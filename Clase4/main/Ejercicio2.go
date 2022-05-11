/*
Structs
Crears una estructura e instanciar ejemplos. Las structs deben poder:
● Almacenar los datos de presentes en el DNI de una persona.
● Representar el perfil bancario de una persona: Nombre y apellido,
dirección, nota de historial crediticio y número de cuenta. El número
de cuenta debe ser privado.
*/

package main

import "fmt"

type DNI struct {
	Name, LastName, DocNacIde string
	BirthDay                  int
	BirthMonth                string
	BirthYear                 int
	Address                   string
}

type BankProfile struct {
	DNI
	/* El Banco Central de la República Argentina asigna un puntaje crediticio
	del 1 al 6 para cada perfil en función de la información sobre su comportamiento.*/
	CreditHistory     int
	bankAccountNumber int
}

func (s *BankProfile) show() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Last Name: ", s.LastName)
	fmt.Println("DNI: ", s.DocNacIde)
	fmt.Println("Birthday: ", s.BirthMonth, "/", s.BirthDay, "/", s.BirthYear)
	fmt.Println("Address: ", s.Address)
	fmt.Println("Credit history: ", s.CreditHistory)
	//fmt.Println("Bank account: ", s.bankAccountNumber, " (pero es un secreto, shhhhh)\n")
}

func main() {
	donPepito := BankProfile{
		DNI: DNI{
			Name:       "Pepito",
			LastName:   "Ibrahimovic",
			DocNacIde:  "19100520",
			BirthDay:   3,
			BirthMonth: "Octubre",
			BirthYear:  1981,
			Address:    "Milán, Italia",
		},
		CreditHistory:     3,
		bankAccountNumber: 452801,
	}
	donPepito.show()

	donJose := BankProfile{
		DNI: DNI{
			Name:       "José",
			LastName:   "Materazzi",
			DocNacIde:  "17565078",
			BirthDay:   19,
			BirthMonth: "Agosto",
			BirthYear:  1973,
			Address:    "Lecce, Italia",
		},
		CreditHistory:     6,
		bankAccountNumber: 075035,
	}
	donJose.show()

}
