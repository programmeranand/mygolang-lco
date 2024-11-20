package main

import "fmt"

func main() {
	fmt.Println("Welcome to function course")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(3, 5, 6, 11)
	fmt.Println("Pro Result is: ", proResult)
	fmt.Println("Pro Result is: ", message)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Sample Message"
}
func greeter() {
	fmt.Println("Nameste from golang")
}

func greeterTwo() {
	fmt.Println("Another Method")
}
