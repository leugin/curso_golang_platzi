package main

import "fmt"

func main() {
	//array
	var array [4]int
	array[0] = 1
	fmt.Println(array, len(array), cap(array))

	// slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// add
	slice = append(slice, 7)

	//add list
	newSlice := []int{11, 12}
	slice = append(slice, newSlice...)

	fmt.Println(slice)

}
