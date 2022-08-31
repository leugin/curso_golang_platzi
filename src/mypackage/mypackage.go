package mypackage

import "fmt"

// Hola se debe correr go mod init para que funcione
type CarPublic struct {
	Brand string // si se coloca en mayuscula es publica si es miniuscula es privad
	year  int
}

func PrintMensaje() {
	fmt.Println("Print")
}
