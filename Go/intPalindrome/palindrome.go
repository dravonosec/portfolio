package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/palindrome-number/description/

func main() {
	num := 121
	fmt.Println(isPalindrome(num))
}

func isPalindrome(x int) bool {
    if (x < 0) {
		return false
	}

	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-i-1] {
            return false
        }
    }

	return true
}