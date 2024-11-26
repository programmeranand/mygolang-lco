package main

import (
	"fmt"
	"log"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcom to maths in Golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4.56

	// fmt.Println("SUm is: ", numberOne+int(numberTwo))
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myRandomNumber)

}
