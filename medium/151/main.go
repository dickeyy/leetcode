package main

import (
	"strings"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.96mb (beats 77.43% of Go solutions)
 */

// Attempt 2
func reverseWords(s string) string {
	words := strings.Fields(s) // splits and removes whitespace
	n := len(words)
	// reverse in place
	for i := range n / 2 {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}
	return strings.Join(words, " ")
}

/*
 *   Stats:
 *   Runtime: 10ms (beats 6.45% of Go solutions)
 *   Memory: 10.01mb (beats 6.11% of Go solutions)
 */

// Attempt 1
// func reverseWords(s string) string {
// 	words := strings.Split(s, " ")
// 	rev := []string{}

// 	for i := range words {
// 		if words[i] != "" {
// 			rev = append([]string{words[i]}, rev...)
// 		}
// 	}

// 	return strings.Join(rev, " ")
// }
