package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in GOLang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Len of Fruit list is ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy List is", len(vegList))
}
