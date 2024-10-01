package main

import "fmt"

const LoginToken string = "gfasoiuwe" //first letter capital means public var

func main() {
	// var username string = "akshat"
	// fmt.Println(username)
	// fmt.Printf("Variable is a type of %T \n", username)
	// var isBoolean bool = true
	// fmt.Println(isBoolean)
	// fmt.Printf("Variable is a type of %T \n", isBoolean)
	// var number uint8 = 255
	// fmt.Println(number)
	// fmt.Printf("Variable is a type of %T \n", number)
	// var number float64 = 255.0923423423423
	// fmt.Println(number)
	// fmt.Printf("Variable is a type of %T \n", number)
	// var number int
	// fmt.Println(number)
	// fmt.Printf("Variable is a type of %T \n", number)

	// var website = "newlink.com"
	// fmt.Println(website)
	// fmt.Printf("Variable is a type of %T \n", website)

	numOfUser := 300000.0 //walrus operator only allowed within method
	fmt.Println(numOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is a type of %T \n", LoginToken)
}