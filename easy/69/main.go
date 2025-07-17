package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.09mb (beats 75.00% of Go solutions)
 */

func mySqrt(x int) int {
	l, r := 0, x
	ans := 0

	for l <= r {
		m := (l + r) / 2
		if m*m <= x {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}
