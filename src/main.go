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
	normalFunction("Hola mundo")
	tripeAdmin(1, 2, "3")

	value := returnValue(5)
	fmt.Println(value)

	value1, value2 := returnValue2(2)
	fmt.Println(value1, value2)

	// tomando solo e primer val
	value3, _ := returnValue2(2)
	fmt.Println(value3)
}
