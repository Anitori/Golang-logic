package main

import "fmt"

func main (){
fmt.Println(uno(5))

numero, estado := dos(1)
fmt.Println(numero)
fmt.Println(estado)

fmt.Println(tres(2))

fmt.Println((cuatro(5,46)))
fmt.Println((cuatro(1,10)))
fmt.Println((cuatro(6,4)))
fmt.Println((cuatro(2,2,5,30,64)))

fmt.Println((cinco(1,1)))
fmt.Println((cinco(2,2)))
fmt.Println((cinco(3,3)))
fmt.Println((cinco(1,1,1,1,50)))

}

////////////////////////////////////////

func uno(numero int) int{
	return numero*2
}

////////////////////////////////////////

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

////////////////////////////////////////

func tres(numero int) (z int){
	z = numero*2
	return
}

// PARAMETROS VARIABLES
// Parametros dinámicos, donde yo no se el número de parámetros de entrada
// Se colocan tres puntos adelante del tipo de dato
// La instrucción range me devuelve 2 valores
// Se puede usar el _ para alojar un tipo de dato, aunque no lo quiera utilizar

func cuatro(numero ...int) int{
	total := 0
	for _, num := range numero {
		total = total + num

	}
	return total

}

/////////////////////

func cinco(numero ...int) int{
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)

	}
	return total

}
// Resultado de Función cinco

//Item 0
// Item 1
//2
//Item 0
//Item 1
//4
// Item 0
//Item 1
// 6
// Item 0
// Item 1
//Item 2
//Item 3
// Item 4
//54