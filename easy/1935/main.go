package main

import "strings"

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.38mb (beats 8.31% of Go solutions)
 */

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	chars := strings.Split(brokenLetters, "")
	count := 0

	for _, word := range words {
		for _, char := range chars {
			if strings.Contains(word, char) {
				count++
				break
			}
		}
	}

	return len(words) - count
}
