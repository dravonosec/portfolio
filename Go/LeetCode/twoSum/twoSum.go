package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		mapIndex, exist := numMap[diff]
		if exist {
			return []int{mapIndex, i}
		}
		numMap[num] = i
	}

	return []int{0, 0}
}
