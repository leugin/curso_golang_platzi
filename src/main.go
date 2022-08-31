package main

import "fmt"

type pc struct {
	brand string
	ram   int
	disk  int
}

func (myPc pc) ping() {

	fmt.Println("brand ", myPc.brand, myPc.ram)
}

func (myPc *pc) dobleRam() {

	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a // &se accede a la direccion de memoria
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // se accede al valor de la direccion

	*b = 110 // si modificamos *b

	fmt.Println(a) // se modifica en el original tambien

	myPc := pc{brand: "adf", disk: 5, ram: 50}
	myPc.ping()

	myPc.dobleRam()

	myPc.ping()
}
