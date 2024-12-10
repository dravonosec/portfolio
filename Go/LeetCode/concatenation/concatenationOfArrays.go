package main

import "fmt"

// https://leetcode.com/problems/concatenation-of-array/description/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) [] int {
	result := make([]int, len(nums)*2)
	for i, value := range nums {
		result[i] = value
		result[i + len(nums)] = value
	}

	return result
}
