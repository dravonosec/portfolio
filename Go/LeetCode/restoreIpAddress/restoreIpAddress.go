package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/restore-ip-addresses/description/

func main() {
	// s := "25525511135" // ["255.255.11.135","255.255.111.35"]
	s := "101023" // "["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}

	if len(s) == 4 {
		return []string{switchPoints(s, 1, 2, 3)}
	}

	result := []string{}
	for i := 1; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			for k := j + 1; k < len(s)-0; k++ {
				if isValid(s[:i]) && isValid(s[i:j]) && isValid(s[j:k]) && isValid(s[k:]) {
					result = append(result, switchPoints(s, i, j, k))
				}
			}
		}
	}

	return result
}

func switchPoints(s string, i int, j int, k int) string {
	builder := strings.Builder{}
	builder.WriteString(s[:i])
	builder.WriteString(".")
	builder.WriteString(s[i:j])
	builder.WriteString(".")
	builder.WriteString(s[j:k])
	builder.WriteString(".")
	builder.WriteString(s[k:])
	return builder.String()
}

func isValid(s string) bool {
	if len(s) > 3 {
		return false
	}

	val, err := strconv.Atoi(s)
	if err != nil || (strings.HasPrefix(s, "0") && len(s) > 1) {
		return false
	}

	if val >= 0 && val <= 255 {
		return true
	} else {
		return false
	}
}
