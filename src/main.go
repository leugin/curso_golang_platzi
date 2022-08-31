package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["hola"] = 15
	m["hh"] = 16
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	// obtner valor

	value, ok := m["hola"]
	fmt.Println(value, ok)

	// obtner validar si existeS

	value2, ok2 := m["hola3"]
	fmt.Println(value2, ok2)
}
