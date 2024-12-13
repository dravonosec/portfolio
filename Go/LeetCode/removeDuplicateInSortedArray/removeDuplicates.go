package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
