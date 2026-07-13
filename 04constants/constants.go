package main

import "fmt"

const val = 10

var txt = "hello"

func main() {
	const name string = "golang"
	// name = "css"
	const age = 21

	fmt.Println(val)
	fmt.Println(txt)

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(host, port)
}