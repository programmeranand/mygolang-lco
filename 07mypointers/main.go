package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is", ptr)

	myNumber := 23

	// & Here means Reference
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)

	// * is used to deferecne the memory address
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is ", myNumber)

}
