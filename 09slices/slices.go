package main

import (
	"fmt"
	"slices"
)

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
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
	
	numbers[0] = 4
	numbers[1] = 5
	fmt.Println(numbers)

	var nums1 = make([]int, 0, 5)
	nums1 = append(nums1, 2)
	var nums2 = make([]int, len(nums1))
	copy(nums2, nums1)
	fmt.Println(nums1, nums2)

	var slc1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slc1[1:4])
	fmt.Println(slc1[:2])
	fmt.Println(slc1[1:])

	var slc2 = []int{1, 2, 5}
	var slc3 = []int{1, 2, 0}
	fmt.Println(slices.Equal(slc2, slc3))

	var slc4 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(slc4)
}