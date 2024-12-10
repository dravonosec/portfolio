package main

// https://leetcode.com/problems/longest-common-prefix/description/

import (
	"fmt"
)

func main() {
	test := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(test)
}

func longestCommonPrefix(strs []string) string {
	
	var prefix string = "" 
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i]!=strs[0][i] {
                return prefix
            }
		}
		prefix = prefix + string(strs[0][i])
	}
	
	return prefix
}