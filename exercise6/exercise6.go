package main

import "fmt"

func main() {
	/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	Por otro lado también es necesario:
	- Saber cuántos de sus empleados son mayores a 21 años.
	- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	- Eliminar a Pedro del mapa.*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Los empleados mayores de 21 años son:%-10s \n\n", "")

	for k, v := range employees {

		if v >= 21 {

			fmt.Printf("%-10s %d \n\n", k, v)

		}

	}

	employees["Federico"] = 25
	fmt.Printf("Se ha Agregado a Federico a la lista%-10s\n\n", "")

	delete(employees, "Pedro")
	fmt.Printf("Se ha eliminado a Pedro de la lista%-10s\n\n", "")

	fmt.Printf("La lista de empleados es :%s \n\n", "")
	for k, v := range employees {

		fmt.Printf("%-10s %d \n\n", k, v)

	}
}
