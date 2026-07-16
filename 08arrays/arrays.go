package main

import "fmt"

func main() {
	// fixed size, predictable, memory optimization, constant time access
	var nums [4]int
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	var vals [4]bool
	fmt.Println(vals)

	var strs [4]string 
	fmt.Println(strs)

	numArr := [3]int{1, 2, 3}
	fmt.Println(numArr)

	nums2d := [2][2]int{{1, 2}, {5, 6}}
	fmt.Println(nums2d)
}