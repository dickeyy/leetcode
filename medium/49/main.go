package main

import (
	"slices"
)

/*
 *   Stats:
 *   Runtime: 3ms (beats 95.56% of Go solutions)
 *   Memory: 9.02mb (beats 89.33% of Go solutions)
 */

func groupAnagrams(strs []string) [][]string {
	h := map[string][]string{}
	for _, s := range strs {
		ca := []byte(s)
		slices.Sort(ca)
		key := string(ca)
		h[key] = append(h[key], s)
	}

	out := make([][]string, 0, len(h))
	for _, group := range h {
		out = append(out, group)
	}

	return out
}
