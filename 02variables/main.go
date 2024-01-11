package main

import "fmt"

const (
	LoginToken string = "gibberish"
)

func main() {
	var username string = "Saurabh"

	fmt.Println(username)
	fmt.Printf("Type of variable: %T\n", username)

	var isLoggedIn bool = true

	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable: %T\n", isLoggedIn)

	var bigFloat float64 = 155.979287491874

	fmt.Println(bigFloat)
	fmt.Printf("Type of variable: %T\n", bigFloat)

	withoutVar := "without var"
	fmt.Println(withoutVar)
	fmt.Printf("Type of variable: %T\n", withoutVar)

	fmt.Println(LoginToken)
	fmt.Printf("Type of variable: %T\n", LoginToken)

}
