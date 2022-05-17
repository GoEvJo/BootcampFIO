/*
Structs y métodos
Crear un sistema que permita gestionar cursos.
● La entidad Curso debe contar con: ID, nombre, aula asignada y
cantidad máxima de estudiantes.
● Debe contar con una lista de estudiantes y permitir su listado, alta,
baja y modificación.
● Debe contar con un profesor asignado: ID, nombre, antigüedad.
● Cada alumno debe contar con: ID, nombre, lista de notas y cantidad
de faltas.
● Crear un método que permita listar y devolver qué alumnos se
encuentran por debajo de 6 en su promedio de notas.
● Crear un método que permita listar y devolver qué alumnos están
libres por cantidad de faltas, siendo hasta 3 faltas las permitidas.
● Crear un método que permita listar y devolver los 3 alumnos con
mejor promedio.
● Crear un método que devuelva cuán llena en porcentaje está el aula
y cuántos alumnos tiene asignados.
*/

package main

import "fmt"

type Identity struct {
	ID     int
	nombre string
}
type alumno struct {
	Identity
	notas  [3]int
	faltas int
}

func (alumno alumno) prom(al alumno) float64 {
	promedio := (float64(al.notas[0]) + float64(al.notas[1]) + float64(al.notas[2])) / 3
	return promedio
}

type profesor struct {
	Identity
	antigüedad int
}
type Curso struct {
	Identity
	aulaAsignada, capMax int
	lista                []alumno
	encargado            profesor
}

func (metodo *Curso) Reprobados(laLista []alumno) {
	// Recibo la lista de alumnos, tengo que acceder a las notas de cada alumno, sacar el promedio y despues comparar con 6
	for i, _ := range laLista {
		fmt.Println("rangos de slice")
		if laLista[i].prom(laLista[i]) < 6 {
			println("Reprobado")
		}
	}
}
func (metodo *Curso) Libres(laLista []alumno) {

}
func (metodo *Curso) Mejorcitos(laLista []alumno) {

}
func (metodo *Curso) Ocupacion(laLista []alumno) {

}

func main() {
	fmt.Println("No me borres el import fmt porfa")

	listaDeAlumnos := []alumno{}

	profesor1 := profesor{
		Identity: Identity{
			ID:     1505,
			nombre: "Santiago",
		},
		antigüedad: 15,
	}

	alumno1 := alumno{
		Identity: Identity{
			ID:     1910,
			nombre: "José",
		},
		notas:  [3]int{5, 9, 6},
		faltas: 2,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno1)

	alumno2 := alumno{
		Identity: Identity{
			ID:     1912,
			nombre: "Cintia",
		},
		notas:  [3]int{6, 10, 7},
		faltas: 3,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno2)

	alumno3 := alumno{
		Identity: Identity{
			ID:     1913,
			nombre: "Simón",
		},
		notas:  [3]int{6, 6, 6},
		faltas: 0,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno3)

	alumno4 := alumno{
		Identity: Identity{
			ID:     1914,
			nombre: "José",
		},
		notas:  [3]int{4, 5, 2},
		faltas: 5,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno4)

	alumno5 := alumno{
		Identity: Identity{
			ID:     1915,
			nombre: "Ana",
		},
		notas:  [3]int{10, 9, 8},
		faltas: 1,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno5)

	alumno6 := alumno{
		Identity: Identity{
			ID:     1910,
			nombre: "Rolando",
		},
		notas:  [3]int{10, 8, 7},
		faltas: 0,
	}
	listaDeAlumnos = append(listaDeAlumnos, alumno6)

	Curso1 := Curso{
		Identity: Identity{
			ID:     1808,
			nombre: "Física cuántica avanzada",
		},
		aulaAsignada: 1,
		capMax:       10,
		lista:        listaDeAlumnos,
		encargado:    profesor1}
}
