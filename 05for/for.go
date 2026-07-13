package main

import "fmt"

// for loop is the only looping construct in go
func main() {
	// while loop implementation using "for"
	i := 1
	for i <= 3{
		fmt.Println(i)
		i += 1
	}

	// // infinite loop
	// for {
	// 	println("!")
	// }

	// classic for loop
	for  i := 0; i <= 3; i++ {
		// break
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// range (3 is excluded)
	for i := range 11 {
		fmt.Println(i)
	}
}