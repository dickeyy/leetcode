package main

import (
	"math/big"
)

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.44mb (beats 47.25% of Go solutions)
 */

func addBinary(a string, b string) string {
	aInt := new(big.Int)
	bInt := new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)

	return aInt.Add(aInt, bInt).Text(2)
}
