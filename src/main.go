package main

import "fmt"

func main() {

	switch modulo := 2 % 2; modulo {
	case 0:
		fmt.Println("es pard")
	default:
		fmt.Println("es impard")
	}

	value := 200
	// Sin condicion
	switch {
	case value > 100:
		fmt.Println("MAyor a 100")
	case value < 0:
		fmt.Println("minor a 0")
	default:
		fmt.Println("Otro")
	}
}
