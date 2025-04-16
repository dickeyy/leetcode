package main

import (
	"strconv"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.90mb (beats 37.89% of Go solutions)
 */

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	writeIndex := 0
	currentChar := chars[0]
	count := 1

	for i := 1; i < len(chars); i++ {
		if chars[i] == currentChar {
			count++
		} else {
			chars[writeIndex] = currentChar
			writeIndex++

			if count > 1 {
				countStr := []byte(strconv.Itoa(count))
				for _, digit := range countStr {
					chars[writeIndex] = digit
					writeIndex++
				}
			}

			currentChar = chars[i]
			count = 1
		}
	}

	chars[writeIndex] = currentChar
	writeIndex++

	if count > 1 {
		countStr := []byte(strconv.Itoa(count))
		for _, digit := range countStr {
			chars[writeIndex] = digit
			writeIndex++
		}
	}

	return writeIndex
}
