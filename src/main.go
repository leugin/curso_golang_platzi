package main

import "fmt"

func main() {
	hello := "Hola"
	fmt.Println(hello)
	nombre := "platzi"
	cursos := 500

	//printf
	fmt.Printf("%s hola tiene mas de %d\n", nombre, cursos)
	fmt.Printf("%v hola tiene mas de %v\n", nombre, cursos)

	//Spring
	message := fmt.Sprintf("%s hola tiene mas de %d\n", nombre, cursos)

	fmt.Println(message)

	// tipo de datos
	fmt.Printf("Hello %T\n", message)
}
