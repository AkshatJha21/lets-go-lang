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

	var numsArr = make([]int, 0, 5)
	fmt.Println(numsArr)
	fmt.Println(numsArr == nil)
	// cap => capacity => maxm number of elements can fit
	fmt.Println(cap(numsArr))
	numsArr = append(numsArr, 1)
	numsArr = append(numsArr, 2)
	numsArr = append(numsArr, 3)
	numsArr = append(numsArr, 4)
	numsArr = append(numsArr, 5)
	numsArr = append(numsArr, 6)
	numsArr = append(numsArr, 7)
	numsArr = append(numsArr, 8)
	numsArr = append(numsArr, 9)
	numsArr = append(numsArr, 10)
	numsArr = append(numsArr, 11)
	fmt.Println(numsArr)
	fmt.Println(cap(numsArr))

	numbers := []int{}
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	fmt.Println(numbers);
	fmt.Println(cap(numbers));
	fmt.Println(len(numbers));
}