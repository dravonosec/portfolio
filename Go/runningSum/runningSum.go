package main

import "fmt"

// https://leetcode.com/problems/running-sum-of-1d-array/
func main() {
	nums := []int{1, 2, 3, 4, 5}
    fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
    result := make([]int,len(nums))

	for i, val := range nums {
		if i == 0 {
			result[i] = val
		} else {
			result[i] = val + result[i-1]
		}
	}

	return result
}