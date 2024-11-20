package main

import "fmt"

func main() {

	fmt.Println("if else in golang")
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "power user"
	} else {
		result = "Exact 10 login counts"
	}
	fmt.Println(result)

	// Check odd even
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

}