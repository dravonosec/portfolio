package main

import "fmt"

// https://leetcode.com/problems/isomorphic-strings/description/

func main() {
	s := "egg"
	t := "add"

	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	// По условию гарантируется, что строки одной длины
	// if (len(s) != len(t)) {
	// 	return false
	// }

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		valS, exist := sToT[charS]
		if exist && charT != valS {
			return false
		} else {
			sToT[charS] = charT
		}

		valT, exist := tToS[charT]
		if exist && charS != valT {
			return false
		} else {
			tToS[charT] = charS
		}
	}

	return true
}
