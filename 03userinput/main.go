package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "dqqdq"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the Pizza")

	// comma ok  || err ok
	// GO does not have a try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of Rating is %T", input)
}
