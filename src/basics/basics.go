package basics

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Examples(){
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

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El area del cuadrado es:", areaCuadrado)

	// Area triangulo rectangulo
	cateto1 := 12
	cateto2 := 26

	areaTrianguloRectangulo := areaTrianguleR(cateto1, cateto2)
	fmt.Println("Area triangulo rectangulo:", areaTrianguloRectangulo)

	// Area circulo
	var diametro float64 = 25
	var radio float64 = diametro/2

	areaCirculo := areaCircle( radio, diametro)
	fmt.Println("Area circulo:", areaCirculo)

	// Area trapecio
	base1 := 3
	base2 := 5
	altura := 7

	areaTrapecio := areaTrapecio(base1, base2, altura)
	println("Area trapecio:", areaTrapecio)
}

func areaTrianguleR(c1, c2 int) int{
	var area int = (c1 * c2) / 2

	return area
}

func areaCircle(r, d float64) float64{
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

func FmtPackage(){
	hello := "helloworld"
	number := 500

	message := fmt.Sprintf("%s %d", hello, number)

	fmt.Printf("%s type: %T \n", message, message)
}

func For(lista []int){
	for i, par := range lista {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

func Caster(s string, i int) (string, int){
	var sToNum int 
	var err error
	var iToStr string

	if s != ""{
		sToNum, err = strconv.Atoi(s)
		if err != nil{
			log.Fatalln(err)
		}
	}
	 
	iToStr = strconv.Itoa(i) 

	fmt.Println(iToStr, sToNum)
	return iToStr, sToNum
}

func IsPair(s string, i int) bool{
	var sToNum int 
	var err error

	if s != "" {
		sToNum, err = strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		
		if sToNum%2 != 0 {
			return false
		} else {
			return true
		}
	}

	if i%2 != 0 {
		return false
	} else {
		return true
	}
}

func Day(i uint) error{
	switch i{
	case 1:
		fmt.Println("Es Lunes")
		return nil
	case 2:
		fmt.Println("Es Martes")
		return nil
	case 3:
		fmt.Println("Es Miercoles")
		return nil
	case 4:
		fmt.Println("Es Jueves")
		return nil
	case 5:
		fmt.Println("Es Viernes")
		return nil
	case 6:
		fmt.Println("Es Sabado")
		return nil
	case 7:
		fmt.Println("Es Domingo")
		return nil
	default:
		return fmt.Errorf("el numero %d no corresponde a ningun dia, valores aceptados entre 1 y 7", i)
	}
}

func IsPalindrome(s string) bool {
	var result string
	sL := strings.ToLower(s)

	for _, letter := range sL {
        result = string(letter) + result
    }

	fmt.Println(s, result)
	return sL == result
}

func Maps(){
	m := make(map[string]int)

	m["mercury"] = 1
	m["venus"] = 2
	m["earth"] = 3
	m["mars"] = 4

	value, ok := m["jupiter"]
	fmt.Println(value, ok)
}

type Car struct {
	Brand string
	Year uint16
	Model string
	SubClass anotherStruct
}

type anotherStruct struct {
	val string
}

func StructingBabyyyy (){
	ferrari := Car{Brand: "Ferrari",
					Year: 2020,
					Model: "Roma",
					SubClass: anotherStruct{val: "coupe",}}
					

	ford := Car{Brand: "Ford"}

	fmt.Println(ferrari)
	fmt.Println(ford)
}

func (myCar Car) Ping(){

	fmt.Println(myCar.Brand, "pong")
}

func (myCar *Car) ChangeBrand(brand string){
	myCar.Brand = brand
	fmt.Println("The brand has changed to: " + myCar.Brand)
}

func removeNegatives(A []int) []int{

	slice := make([]int, 0)

	for _, val := range A{
		
		if A[len(A)-1] < 0{
			
			return slice
		}
		if val > 0 {

		slice =	append(slice, val)
		}
	}

	return slice
}

func Solution(A []int) int {

    sort.Ints(A)
	positives := removeNegatives(A)

	last := 0
	
    for _, value := range positives{

		if value != last && last+1 != value{

			return last+1
		}

		last = value

	}

	return last+1
}