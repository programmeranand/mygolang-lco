package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.co.in"

func main() {

	fmt.Println("Welcome to Web request turorial")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type %T\n", response)

	// perform close at the end, perform rest of the operations post this
	defer response.Body.Close() // Callers responsibiility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// typcast databytes to string
	content := string(databytes)
	fmt.Println("Response content is: ", content)

}
