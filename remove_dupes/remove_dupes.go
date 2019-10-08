// Given a sorted array nums, remove the duplicates in-place such that
// each element appears only once and return the new length.
// Do not allocate extra space for another array,
// you must do this by modifying the input array in-place with O(1) extra memory.

package main

import "fmt"

func removeDuplicates(nums []int) int {
	fmt.Println(nums)
	return len(nums)
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 5, 6, 6}
	result := removeDuplicates(nums)
	fmt.Println(result)
}
