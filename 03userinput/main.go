package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate my program: ")

	// comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of rating: %T", input)
}