package main

import (
	"strings"

	"github.com/dickeyy/leetcode/utils"
)

/*
 *   Stats:
 *   Runtime: 74ms (beats 5.71% of Go solutions)
 *   Memory: 30.15mb (beats 9.76% of Go solutions)
 */

func main() {
	utils.AssertEqual("Hi Im Here", addSpaces("HiImHere", []int{2, 4, 8}), 1)
}

func addSpaces(s string, spaces []int) string {
	spaceMap := make(map[int]bool)
	for _, pos := range spaces {
		spaceMap[pos] = true
	}

	builder := strings.Builder{}
	builder.Grow(len(s) + len(spaces))

	for i := 0; i < len(s); i++ {
		if spaceMap[i] {
			builder.WriteByte(' ')
		}
		builder.WriteByte(s[i])
	}

	return builder.String()
}