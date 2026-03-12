package main

/*
 *   Stats:
 *   Runtime: 8ms (beats 36.76% of Go solutions)
 *   Memory: 8.23mb (beats 49.71% of Go solutions)
 */

func minFlips(s string) int {
	n := len(s)
	out := n
	ss := s + s

	pos := [][]int{{0, 0}, {0, 0}}

	// compute initial pattern
	for j := 0; j <= n-1; j++ {
		if j%2 == 0 {
			if ss[j] == '0' {
				pos[0][0]++
			} else {
				pos[1][0]++
			}
		} else {
			if ss[j] == '0' {
				pos[0][1]++
			} else {
				pos[1][1]++
			}
		}
	}

	out = min(out, min(pos[1][0]+pos[0][1], pos[0][0]+pos[1][1]))

	// slide n-1 times
	for i := 0; i <= n-2; i++ {
		if ss[i] == '0' {
			pos[0][0]--
		} else {
			pos[1][0]--
		}

		// swap evens and odds
		og0, og1 := pos[0][0], pos[1][0]
		pos[0][0], pos[1][0] = pos[0][1], pos[1][1]
		pos[0][1], pos[1][1] = og0, og1

		// compute new patterns
		if (n-1)%2 == 0 {
			if ss[i+n] == '0' {
				pos[0][0]++
			} else {
				pos[1][0]++
			}
		} else {
			if ss[i+n] == '0' {
				pos[0][1]++
			} else {
				pos[1][1]++
			}
		}

		out = min(out, min(pos[1][0]+pos[0][1], pos[0][0]+pos[1][1]))
	}

	return out
}
