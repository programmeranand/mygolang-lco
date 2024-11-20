package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano()) //random seed
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice val is 1, you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again")
	}
}
