package main

import "math"

/*
 *   Stats:
 *   Runtime: 0ms (beats 100% of Go solutions)
 *   Memory: 3.81mb (beats 60.00% of Go solutions)
 */

// Could reduce memory usage by creating a
// intPow function to avoid casting to floats

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// Find minimum x such that (T+1)^x >= N
	t := math.Floor(float64(minutesToTest) / float64(minutesToDie))
	x := 0
	for math.Pow(t+1, float64(x)) < float64(buckets) {
		x++
	}

	return x
}
