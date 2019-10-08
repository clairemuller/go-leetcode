// https://leetcode.com/problems/two-sum/
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

package main

import "fmt"

// nested for loops
func twoSum(nums []int, sum int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == sum {
				sum := []int{i, j}
				return sum
			}
		}
	}
	return []int{}
}

// using a map
func twoSumMap(nums []int, sum int) []int {
	// initialize empty map
	m := make(map[int]int)

	// start looping thru nums
	// target is the current number's match to fulfill sum
	for i := 0; i < len(nums); i++ {
		target := sum - nums[i]
		// if the map has a key of target
		// return the target's index and current index
		// else create the key/value pair
		if _, ok := m[target]; ok {
			return []int{m[target], i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 11, 15, 7}
	sum := 9
	result := twoSumMap(nums, sum)
	fmt.Println(result)
}
