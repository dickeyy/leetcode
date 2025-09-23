package main

import (
	"strconv"
	"strings"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 3.82mb (beats 96.80% of Go solutions)
 */

func compareVersion(version1 string, version2 string) int {
	vs1 := strings.Split(version1, ".")
	vs2 := strings.Split(version2, ".")

	maxLen := max(len(vs2), len(vs1))

	for i := range maxLen {
		vm1 := 0
		if i < len(vs1) {
			vm1, _ = strconv.Atoi(vs1[i])
		}

		vm2 := 0
		if i < len(vs2) {
			vm2, _ = strconv.Atoi(vs2[i])
		}

		if vm1 > vm2 {
			return 1
		} else if vm1 < vm2 {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
