package main

import "unicode"

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.10mb (beats 43.30% of Go solutions)
 */

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	var hasVowel, hasConsonant bool
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	for _, char := range word {
		if !unicode.IsDigit(char) && !unicode.IsLetter(char) {
			return false
		}

		if unicode.IsLetter(char) {
			if vowels[char] {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		}
	}

	return hasVowel && hasConsonant
}
