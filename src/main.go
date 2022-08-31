package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}
func main() {
	c := make(chan string, 2)
	c <- "Mensaje"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c))

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	/// select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Hola", email1)
	go message("Adios", email2)
	for i := 0; i < 2; i++ {
		select {
		case n1 := <-email1:
			fmt.Println("Respondio 1: ", n1)
		case n2 := <-email2:
			fmt.Println("respondio 2", n2)
		}
	}

}
