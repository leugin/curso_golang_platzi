package main

import "fmt"

type figurad2D interface {
	area() float64
}
type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func calcular(fig figurad2D) {
	fmt.Println(fig.area())
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (c rectangulo) area() float64 {
	return c.base * c.altura
}
func main() {
	myCuadrado := cuadrado{base: 5}
	myRectanguls := rectangulo{base: 5, altura: 2}

	fmt.Println(myCuadrado.area())
	fmt.Println(myRectanguls.area())

	// por interfaz
	calcular(myCuadrado)
	calcular(myRectanguls)

	/// lita de interfaces

	myInterfaes := []interface{}{"Hola", 5, 20}

	fmt.Println(myInterfaes...)
}
