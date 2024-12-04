package main

import "github.com/dickeyy/leetcode/utils"

func main() {
	utils.AssertEqual(0, strStr("sadbutsad", "sad"), 1)
}

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 3.93mb (beats 6.30% of Go solutions)
 */

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if len(needle)+i > len(haystack) {
			return -1
		} else if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
