package main

import "fmt"

// Variable Name if Starts with Capital letter means it is publicicaly avaiialbe
// here LoginToken Constant is publicly available
const LoginToken string = "dqdqdqd"

func main() {
	var username string = "Ankit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFLoat float32 = 255.4523232331121
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T \n", smallFLoat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learcodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
