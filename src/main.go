package main

import "fmt"

func isPalindromo(text string) {
	var textCopy string
	for i := len(text) - 1; i >= 0; i-- {
		textCopy += string(text[i])
	}
	if text == textCopy {
		fmt.Println("is paili")
	} else {
		fmt.Println("no es pali")
	}
}
func main() {

	slice := []string{"Hola", "Mundo"}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("hola")
	isPalindromo("ama")
}
