package main

import "strings"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 3.85mb (beats 92.19% of Go solutions)
 */

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, word := range strings.Split(sentence, " ") {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}

	return -1
}
