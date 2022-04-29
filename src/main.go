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

	// For
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	basics.For(listaNumerosPares)

	// If practice and errors
	basics.Caster("", 256)

	r := basics.IsPair("", 3)
	fmt.Println(r)

	// Switch
	err := basics.Day(1)
	if err != nil{
		fmt.Println(err)
	}

	// Slices
	// var text string
	// fmt.Scan(&text)
	// basics.IsPalindrome(text)

	// Map
	basics.Maps()

	// Structs
	basics.StructingBabyyyy()

	myCar := basics.Car{Brand: "porsche", Model: "911" }

	myCar.Ping()
	myCar.ChangeBrand("BMW")


	a := []int{1,2,3}
	num := basics.Solution(a)
	fmt.Println(num)



	f := []int{0,0,1,1,2}
	n := basics.Lonelyinteger(f)
	fmt.Println(n)
}
