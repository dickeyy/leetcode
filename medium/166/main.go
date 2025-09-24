package main

import (
	"fmt"
	"strings"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.10mb (beats 83.33% of Go solutions)
 */

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := &strings.Builder{}
	if (numerator < 0) != (denominator < 0) {
		res.WriteString("-")
	}

	num := abs(int64(numerator))
	den := abs(int64(denominator))

	fmt.Fprintf(res, "%d", num/den)
	num %= den

	if num == 0 {
		return res.String()
	}

	res.WriteString(".")
	remainderMap := make(map[int64]int)

	for num != 0 {
		if idx, ok := remainderMap[num]; ok {
			s := res.String()
			return s[:idx] + "(" + s[idx:] + ")"
		}
		remainderMap[num] = res.Len()
		num *= 10
		fmt.Fprintf(res, "%d", num/den)
		num %= den
	}

	return res.String()
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
