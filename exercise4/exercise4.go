package main

import "fmt"

func main() {

	/*
		Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
		que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
		por qué?
		Ej: 7, Julio
	*/

	mapex := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	for k, v := range mapex {

		if k == 7 {
			fmt.Printf("%d %s %s", k, ",", v)
		}

	}

}
