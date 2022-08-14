package main

//Puede llegar a tener utilidad en backoffice

import (
	"fmt"
	"os"
	"bufio" //Alternativa de aceptar cosas por teclado
)

var numero1 int
var numero2 int

func main() {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanln(&numero1) // &puntero

	fmt.Println("Ingrese número 2: ")
	fmt.Scanln(&numero2)
}
