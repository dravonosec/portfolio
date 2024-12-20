package main

// https://leetcode.com/problems/divide-two-integers/description/
// Time Limit :(

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, -1))
}

func divide(dividend int, divisor int) int {
	isPositive := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)

	absDevisor := math.Abs(float64(divisor))
	absDividend := math.Abs(float64(dividend))
	if absDividend < absDevisor {
		return 0
	}

	if absDevisor == absDividend {
		return getSignedValue(1, isPositive)
	}

	result := 0

	upperBound := math.Pow(2, 31) - 1
	lowerBound := -math.Pow(2, 31)

	for int(absDividend) >= int(absDevisor) {
		absDividend -= absDevisor
		if isPositive {
			result++
		} else {
			result--
		}

		if (result >  int(upperBound)) {
			return int(upperBound)
		} else if (result < int(lowerBound)) {
			return int(lowerBound)
		}
	}
	
	return result
}

func getSignedValue(num int, isPositive bool) int {
	if isPositive {
		return num
	} else {
		return -num
	}
}