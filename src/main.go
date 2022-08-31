package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 500}
	fmt.Println(myCar)

	// vacia

	var otherCar car
	otherCar.brand = "Demo"
	fmt.Println(otherCar)
}
