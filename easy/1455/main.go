package main

import (
	"strings"

	"github.com/dickeyy/leetcode/utils"
)

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 3.85mb (beats 92.19% of Go solutions)
 */

func main() {
	utils.AssertEqual(4, isPrefixOfWord("I love eating chips", "chips"), 1)
	utils.AssertEqual(3, isPrefixOfWord("I love eating chips", "love"), 2)
}

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, word := range strings.Split(sentence, " ") {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}

	return -1
}
