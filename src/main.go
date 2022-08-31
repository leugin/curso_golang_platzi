package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Si")
	} else {
		fmt.Println("No")
	}
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Si")
	} else {
		fmt.Println("No")
	}

	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Si")
	} else {
		fmt.Println("No")
	}

	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}
