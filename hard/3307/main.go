package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100% of Go solutions) :)
 *   Memory: 4.76mb (beats 50% of Go solutions) :)
 */

func kthCharacter(k int64, operations []int) byte {
	n := len(operations)
	lens := make([]int64, n+1)
	lens[0] = 1
	for i := 0; i < n; i++ {
		if operations[i] == 0 {
			lens[i+1] = lens[i] * 2
		} else {
			lens[i+1] = lens[i] * 2
		}
		if lens[i+1] >= k {
			lens[i+1] = k
		}
	}

	ch := byte('a')
	for i := n - 1; i >= 0; i-- {
		if k <= lens[i] {
			continue
		}
		k -= lens[i]
		if operations[i] == 1 {
			ch = ((ch-'a'+1)%26 + 'a')
		}
	}

	return ch
}
