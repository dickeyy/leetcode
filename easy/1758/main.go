package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.55mb (beats 10.00% of Go solutions)
 */

func minOperations(s string) int {
	ops := [2]int{0, 0}

	for i, c := range s {
		if i%2 == 0 {
			if c == '0' {
				ops[0]++
			} else {
				ops[1]++
			}
		} else {
			if c == '0' {
				ops[1]++
			} else {
				ops[0]++
			}
		}
	}

	return min(ops[0], ops[1])
}
