package main

import "github.com/dickeyy/leetcode/utils"

func main() {
	utils.AssertEqual(1, finalValueAfterOperations([]string{"X++", "X++", "X--", "++X", "--X"}), 1)
}

func finalValueAfterOperations(operations []string) int {
	x := 0
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "--X":
			x--
		case "X--":
			x--
		case "++X":
			x++
		case "X++":
			x++
		}
	}

	return x
}
