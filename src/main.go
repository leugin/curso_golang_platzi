package main

import "fmt"

type pc struct {
	brand string
	ram   int
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("tengo %d gb", myPc.ram)
}

func main() {
	myPc := pc{brand: "lenovo", disk: 52, ram: 50}

	fmt.Println(myPc)
}
