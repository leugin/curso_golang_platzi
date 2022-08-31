package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeAdmin(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a + 2
}

func returnValue2(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// for condicional

	fmt.Println("CONDICIONAL")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("while")

	/// for while
	counter := 0
	for counter < 10 {
		counter++
		fmt.Println(counter)
	}
	fmt.Println("for ever")

	for {
		fmt.Println(counter)
		counter++
	}
}
