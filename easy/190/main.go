package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.28mb (beats 88.51% of Go solutions)
 */

func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for range 32 {
		ans = ans<<1 | num&1
		num >>= 1
	}

	return ans
}
