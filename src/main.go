package main

import "fmt"

func main() {
	// declaracion de constantes
	const p1 float64 = 3.14
	const p2 float32 = 3.14
	fmt.Println("pi :", p1)
	fmt.Println("p2 :", p2)
	// Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base", base)
	fmt.Println("altura", altura)
	fmt.Println("area", area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areacuadradao := baseCuadrado * baseCuadrado
	fmt.Println(areacuadradao)

}
