package main

import "fmt"

func main() {

	// defer ejecuta l funicon antes de todo muera
	defer fmt.Println("Hola")
	fmt.Println("Mondo")

	// continue
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		if i == 4 {
			break
		}
	}
}
