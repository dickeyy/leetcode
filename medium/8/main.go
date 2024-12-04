package main

import (
	"strconv"
	"strings"

	"github.com/dickeyy/leetcode/utils"
)

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.22b (beats 12.73% of Go solutions)
 */

func main() {
	utils.AssertEqual(42, myAtoi("42"), 1)
	utils.AssertEqual(-42, myAtoi("  -42"), 2)
	utils.AssertEqual(1337, myAtoi("1337c0d3"), 3)
	utils.AssertEqual(-2147483648, myAtoi("-91283472332"), 4)
	utils.AssertEqual(-5, myAtoi("-5-"), 5)
}

func myAtoi(s string) int {
	rs := strings.Builder{}
	rs.Grow(len(s))

	s = strings.TrimSpace(s)

	for i, c := range s {
		if c == '-' || c == '+' {
			if i == 0 {
				rs.WriteRune(c)
				continue
			}
		}

		if c >= '0' && c <= '9' {
			rs.WriteRune(c)
		} else {
			break
		}
	}

	i, _ := strconv.Atoi(rs.String())
	if i > 2147483647 {
		return 2147483647
	} else if i < -2147483648 {
		return -2147483648
	}

	return i
}
