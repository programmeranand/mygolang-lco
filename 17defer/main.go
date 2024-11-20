package main

import "fmt"

func main() {
	defer fmt.Println("Welcome to Defer course")
	defer fmt.Println("Welcome to Defer course 2")
	defer fmt.Println("Welcome to Defer course 3")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
