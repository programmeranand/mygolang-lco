package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to course on Slices")

	// var fruitList = []string{"apple", "tomato", "peach"}
	// fmt.Printf("Type of of FruitList is %T \n", fruitList)

	// fruitList = append(fruitList, "mango", "Banaa")
	// fmt.Println("FruitList is ", fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 345
	// highScores[2] = 313
	// highScores[3] = 311
	// // highScores[4] = 3233

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	var courses = []string{"reactjs", "python", "java", "c++", "swift"}
	fmt.Println(courses)

	var index int = 2

	// We can use append to remove the values
	courses = append(courses[:index], courses[index+2:]...)
	fmt.Println(courses)
}
