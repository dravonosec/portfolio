package main

import "fmt"

// https://leetcode.com/problems/build-array-from-permutation/description/

func main() {
	numbers := []int{0, 2, 1, 3, 4}
	fmt.Println(buildArray(numbers))
}

func buildArray(nums []int) []int {
    result := make([]int, len(nums))
	for i, value := range nums {
		result[i] = nums[value]
	}

	return result
}