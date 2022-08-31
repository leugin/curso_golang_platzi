package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Hola")
	go say("Mundo", &wg)
	wg.Add(1)
	wg.Wait()
	// funcones anonimas

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)
}
