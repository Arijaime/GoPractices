package main

import "fmt"

func main() {
	/* a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
	necesario generar una aplicación que contenga dicha lista.
	Estudiantes:
	Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
	Leandro, Eduardo, Duvraschka.

	b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
	sin modificar el código que escribiste inicialmente.
	Estudiante:
	Gabriela*/

	var students []string
	students = append(students, "Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "leando", "Eduardo", "Duvraschka")
	fmt.Printf("%10s \n\n", students)
	fmt.Printf("%s\n\n", "--Dos clases despues--")
	students = append(students, "Gabriela")
	fmt.Printf("%10s\n", students)
}
