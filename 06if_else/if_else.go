package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("adult!")
	} else {
		fmt.Println("not adult!")
	}

	umar := 16
	if umar >= 18 {
		fmt.Println("adult!")
	} else if umar >= 12 {
		fmt.Println("teenager!")
	} else {
		fmt.Println("kid!")
	}

	role := "ADMIN"
	hasPermissions := false;

	if role == "ADMIN" && hasPermissions {
		fmt.Println("Yes")
	} else {
		fmt.Println("Access denied")
	}

	if age2 := 15; age2 >= 18 {
		fmt.Println("adult", age2)
	} else if age2 >= 12 {
		fmt.Println("teenage", age2)
	}
}