package main

import "fmt"

//%d - variable de tipo numérico - verbo?
//%s - variable de tipo string - verbo?

var numero int = 40
var texto string
var status bool

// Dentro de los númericos, tenemos más tipos de datos como por ejemplo
// float32 float64
// int8 int32 int64
// uint

//Declarar las variables fuera de las funciones sirve para que se pueda usar en todas

//Para declarar muchas variables se puede hacer
var numero10, numero11, numero12 int

//Otra forma de declarar variables es
//numero2, numero3, numero4 :=2, 5, 67

//Otra forma de declarar variables es por ejemplo
//numero8 := 4

//Varias variables
func main() {
	numero2, numero3, numero4 := 2, 5, 67
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero)

}

//Si la primer letra de una variable esta en minúscula, la detecta como una variable local,
//Si queremos acceder a variables desde otro paquete, tenemos que poner la varaible con una inicial MAYUS

//Si quiero cambiar el tipo de dato de una variable hay que hacer lo siguiente
var moneda2 int = 0
func algo() {
	var moneda int64 = 0

	moneda2 = int(moneda)

}

//Hay un paquete que se llama strconv que sirve para convertir tipos de variable
