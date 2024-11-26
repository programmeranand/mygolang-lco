package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.co.in/news?name=ankit&location=bangalore"

func main() {
	fmt.Println("Welcome to Handling URL's in golang")
	fmt.Println(myurl)

	// parse the URL
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type of query params are: %T\n", qparams)
	fmt.Println(qparams["location"])
	for key, val := range qparams {
		fmt.Println("Key is ", key)
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.co.in",
		Path:    "/news",
		RawPath: "name=ankit",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
