package main 

import "fmt"

//Son funciones sin nombre, se puede modificar en ejecución?
//func como tipo de dato

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}//Devuelve un resultado de tipo entero, variable de tipo función

func main (){
fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5,7))
fmt.Println("Sumo 5 + 7 =", Calculo(5,7))

// Restamos (la gracia de esto es que ahora voy a redifinir "Calculo", lo cual con otra función es imposible)

Calculo = func(num1 int, num2 int) int {
	return num1-num2
}

fmt.Println("Restamos 6 - 4 =", Calculo(6,4))

//Dividimos

Calculo = func(num1 int, num2 int) int {
	return num1/num2
}

fmt.Println("Dividimos 25 / 5 =", Calculo(25,5))

Operaciones()

//Closures

tablaDel := 2
MiTabla := Tabla(tablaDel)
for i := 1; i < 11 ; i++ {
	fmt.Println(MiTabla())
}

}

//Closures

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

////////////////////////////////////////

func Operaciones (){
	resultado := func () int  {

	var a int = 23
	var b int = 13
	return a + b
	}
	fmt.Println(resultado())
}

