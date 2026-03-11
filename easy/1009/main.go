package main

import (
	"strconv"
	"strings"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 3.91mb (beats 15.38% of Go solutions)
 */

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	var out strings.Builder
	str := strconv.FormatInt(int64(n), 2)
	for _, c := range strings.Split(str, "") {
		if c == "1" {
			out.WriteString("0")
		} else {
			out.WriteString("1")
		}
	}
	x, _ := strconv.ParseInt(out.String(), 2, 64)
	return int(x)
}
