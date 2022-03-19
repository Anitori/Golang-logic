package main

import "fmt"

//Se puede declarar una variable en el mismo IF

var estado bool

func main() {
	estado = true
	if estado == true {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	switchCase()
}

func switchCase() {
	numero := 3
	switch numero{
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")	
		
	}
}
