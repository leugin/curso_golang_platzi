package main

import (
	"curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Hola"
	fmt.Println(myCar)

	mypackage.PrintMensaje()
}
