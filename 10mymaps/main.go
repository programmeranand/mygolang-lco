package main

import "fmt"

func main() {

	fmt.Println("Welcome to course on Maps")

	languages := make(map[string]string)

	languages["Js"] = "JavaScript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all Languages", languages)

	// obtaining a value
	fmt.Println("JS Shorts for", languages["Js"])

	//deleting a value
	delete(languages, "rb")
	fmt.Println("List of all Languages", languages)

	// looping in the maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
