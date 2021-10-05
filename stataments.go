
// impresión en consola rápida
package main

func main(){
    print("Hola")
}

//Agregar un alias a un import 

package main

import (
	"fmt"
	mth "math"
)

func main() {
	fmt.Println(mth.Pi)
}

// Zero values de primitivos
var a int // 0
var b float64 // 0
var c string // ""
var d bool // false

// Imprimir tipo de variables
a := 2
fmt.Printf("%T", a)

// Función para tomar los errores

func isError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Ejemplo de uso
func main() {
	_, err := strconv.Atoi("53a")
	isError(err)
}

// Arrays vs Slices

// Array
var myList [2]int

// Slice
var myList2 []int


//Asegurarnos si un key existe en el map
m := make(map[string]int)

m["hola"] = 1

// Nota, usalmente se usa "ok" para recibir la segunda variable
value, ok := m["hello"]

/*
Si existe, ok será "true"
Si no existe, ok será "false"

En este caso, ok es "false" porque no existe.
*/

// Comandos de Go modules

// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod