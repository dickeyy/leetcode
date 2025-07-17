package main

/*
 *   Stats:
 *   Runtime: 70ms (beats 56.22% of Go solutions)
 *   Memory: 14.17mb (beats 69.12% of Go solutions)
 */

func countPrimes(n int) int {
	primes := make([]bool, n)
	count := 0

	for i := 2; i < n; i++ {
		if primes[i] {
			continue
		}

		count++

		for j := i; j < n; j += i {
			primes[j] = true
		}
	}

	return count
}
