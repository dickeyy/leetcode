package main

import (
	"strconv"
	"strings"
)

/*
 *   Stats:
 *   Runtime: 3ms (beats 48.31% of Go solutions)
 *   Memory: 8.73mb (beats 39.74% of Go solutions)
 */

// Now, with the two helper functions, you can start with "1"
// and call the two functions alternatively n-1 times. T
// The answer is the last integer you will obtain.
func countAndSay(n int) string {
	curr := "1"
	for i := 1; i < n; i++ {
		curr = convToStr(mapPairs(curr))
	}

	return curr
}

// helper function that maps an integer to pairs of its digits
// and their frequencies. For example, if you call this function
// with "223314444411", then it maps it to an array of pairs
// [[2,2], [3,2], [1,1], [4,5], [1, 2]]
func mapPairs(str string) [][]int {
	pairs := [][]int{}
	split := strings.Split(str, "")

	for i, num := range split {
		if i == 0 {
			n, _ := strconv.Atoi(num)
			pairs = append(pairs, []int{n, 1})
			continue
		}

		if split[i] == split[i-1] {
			pairs[len(pairs)-1][1]++
		} else {
			n, _ := strconv.Atoi(num)
			pairs = append(pairs, []int{n, 1})
		}
	}

	return pairs
}

// helper function that takes the array of pairs and creates a new integer.
// For example, if you call this function with
// [[2,2], [3,2], [1,1], [4,5], [1, 2]],
// it should create "22"+"23"+"11"+"54"+"21" = "2223115421"
func convToStr(pairs [][]int) string {
	var sb strings.Builder
	for _, pair := range pairs {
		sb.WriteString(strconv.Itoa(pair[1]))
		sb.WriteString(strconv.Itoa(pair[0]))
	}

	return sb.String()
}
