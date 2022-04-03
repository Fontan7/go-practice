package main

import (
	"fmt"
	"math"
)


func main (){
	fmt.Println("Hello world")

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El area del cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x+y
	fmt.Println("Suma:", result)

	// Resta
	result = y-x
	fmt.Println("Resta:", result)

	// Multiplicar
	result = x*y
	fmt.Println("Multiplicacion:", result)

	// Dividir
	result = y/x
	fmt.Println("Dividision:", result)

	// Modulo 
	result = y%x
	fmt.Println("Modulo:", result)

	// Incremental
 	x++ //or assign to a var x + 1 
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Area triangulo rectangulo
	cateto1 := 12
	cateto2 := 26

	areaTrianguloRectangulo := areaTrianguloR(cateto1, cateto2)
	fmt.Println("Area triangulo rectangulo:", areaTrianguloRectangulo)

	// Area circulo
	var diametro float64 = 25
	var radio float64 = diametro/2

	areaCirculo := areaCirculo( radio, diametro)
	fmt.Println("Area circulo:", areaCirculo)

	// Area trapecio
	base1 := 3
	base2 := 5
	altura := 7

	areaTrapecio := areaTrapecio(base1, base2, altura)
	println("Area trapecio:", areaTrapecio)

	// String practice
	fmtPackage()

	// For
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	
}

func areaTrianguloR(c1, c2 int) int{
	var area int = (c1 * c2) / 2

	return area
}

func areaCirculo(r, d float64) float64{
	var pi float64 = 3.14159
	if r > 0 {

		return pi * math.Pow(r, 2)
	}

	return pi * (math.Pow(d, 2) / 4)
}

func areaTrapecio(b1, b2, a int) int{
	var area int = (b1 + b2) * a / 2

	return area
}

func fmtPackage(){
	hello := "helloworld"
	number := 500

	message := fmt.Sprintf("%s %d", hello, number)

	fmt.Printf("%s type: %T \n", message, message)
}

func ifPrac(){
	
}