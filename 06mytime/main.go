package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time Study of GoLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//Chaing the format of the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create time from Values entered
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
