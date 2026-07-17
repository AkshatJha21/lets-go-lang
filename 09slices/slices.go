package main

import "fmt"

func main() {
	// slices = dynamic arrays
	// most used go construct
	// useful arr methods
	var nums []int

	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var numsArr = make([]int, 2)
	fmt.Println(numsArr)
	fmt.Println(numsArr == nil)
	fmt.Println(cap(numsArr))
}