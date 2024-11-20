package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	// Go Lang does not supprt inheritance

	ankit := User{"Ankit", "abc@df.com", true, 29}
	fmt.Println(ankit)
	fmt.Printf("Ankits Details are: %+v\n", ankit)

	// Display specific values from specific
	fmt.Printf("name is %v and Email is: %v\n", ankit.Name, ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
