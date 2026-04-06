package main

import "strings"

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 5.02mb (beats 92.05% of Go solutions)
 */

// func judgeCircle(moves string) bool {
// 	freq := make(map[string]int, 4)
// 	for _, m := range strings.Split(moves, "") {
// 	    freq[m]++
// 	}
// 	return freq["R"] == freq["L"] && freq["U"] == freq["D"]
// }

// func judgeCircle(moves string) bool {
// 	x, y := 0, 0
// 	for m := range strings.SplitSeq(moves, "") {
// 		if m == "U" {
// 			y++
// 		}
// 		if m == "D" {
// 			y--
// 		}
// 		if m == "R" {
// 			x++
// 		}
// 		if m == "L" {
// 			x--
// 		}
// 	}
// 	return x == 0 && y == 0
// }

func judgeCircle(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "R") == strings.Count(moves, "L")
}
