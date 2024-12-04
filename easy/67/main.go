package main

import (
	"math/big"

	"github.com/dickeyy/leetcode/utils"
)

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.44mb (beats 47.25% of Go solutions)
 */

func main() {
	utils.AssertEqual("100", addBinary("11", "1"), 1)
	utils.AssertEqual("10101", addBinary("1010", "1011"), 2)
	utils.AssertEqual("110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		addBinary(
			"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
		),
		3,
	)
}

func addBinary(a string, b string) string {
	aInt := new(big.Int)
	bInt := new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)

	return aInt.Add(aInt, bInt).Text(2)
}
