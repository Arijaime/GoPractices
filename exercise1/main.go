package main

import "fmt"

func main() {
	/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
	depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
	impuesto de un salario.
	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
	sueldo y si gana más de $150.000 se le descontará además un %10.*/

	var sueldo float32 = 120000

	sueldo = sueldo - calculate(sueldo)

	fmt.Println(sueldo)
}

/*	Func that calculate the taxes 	*/
func calculate(salary float32) float32 {

	var taxes float32

	if salary >= 150000 {

		taxes = (salary * 27) / 100

	} else if salary >= 50000 {

		taxes = (salary * 17) / 100

	}

	return taxes
}
