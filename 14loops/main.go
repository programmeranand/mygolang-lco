package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// naive for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// loop index in easy way
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for each loop
	// for index, day := range days {
	// 	fmt.Printf("Index is %v, value is %v\n", index, day)
	// }

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at Another line")
}
