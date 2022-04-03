package main

import (
	"fmt"
	"go-practice/src/basics"
)

func main (){
	fmt.Println("Hello world")

	//Ejemplos operativos
	basics.Examples()

	// String practice
	basics.FmtPackage()

	// If practice
	basics.IfPrac()

	// For
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	basics.For(listaNumerosPares)
	
}
