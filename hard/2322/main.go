package main

import "math"

/*
 *   Stats:
 *   Runtime: 26ms (beats 34.94% of Go solutions) :)
 *   Memory: 6.73mb (beats 59.04% of Go solutions) :(
 */

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)

	// Build adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Root & Precompute: DFS to get subtree XOR + entry/exit times
	parent := make([]int, n)
	subXOR := make([]int, n)
	entryTime := make([]int, n)
	exitTime := make([]int, n)
	timer := 0

	var dfs func(node, par int)
	dfs = func(node, par int) {
		parent[node] = par
		entryTime[node] = timer
		timer++
		subXOR[node] = nums[node]

		for _, child := range adj[node] {
			if child != par {
				dfs(child, node)
				subXOR[node] ^= subXOR[child]
			}
		}

		exitTime[node] = timer
		timer++
	}

	dfs(0, -1)
	totalXOR := subXOR[0]

	result := math.MaxInt32

	// Efficient approach: Try all pairs of nodes as edge endpoints
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			child1, child2 := i, j

			// Case 1: child2 is descendant of child1
			if isInSubtree(child2, child1, entryTime, exitTime) {
				comp1 := subXOR[child2]
				comp2 := subXOR[child1] ^ subXOR[child2]
				comp3 := totalXOR ^ subXOR[child1]
				result = min(result, getScore(comp1, comp2, comp3))
			}
			// Case 2: child1 is descendant of child2
			if isInSubtree(child1, child2, entryTime, exitTime) {
				comp1 := subXOR[child1]
				comp2 := subXOR[child2] ^ subXOR[child1]
				comp3 := totalXOR ^ subXOR[child2]
				result = min(result, getScore(comp1, comp2, comp3))
			}
			// Case 3: child1 and child2 are in different branches
			if !isInSubtree(child1, child2, entryTime, exitTime) && !isInSubtree(child2, child1, entryTime, exitTime) {
				comp1 := subXOR[child1]
				comp2 := subXOR[child2]
				comp3 := totalXOR ^ comp1 ^ comp2
				result = min(result, getScore(comp1, comp2, comp3))
			}
		}
	}

	return result
}

// O(1) check if descendant is in ancestor's subtree using entry/exit times
func isInSubtree(descendant, ancestor int, entryTime, exitTime []int) bool {
	return entryTime[ancestor] <= entryTime[descendant] && exitTime[descendant] <= exitTime[ancestor]
}

func getScore(a, b, c int) int {
	minVal, maxVal := a, a
	if b < minVal {
		minVal = b
	} else if b > maxVal {
		maxVal = b
	}
	if c < minVal {
		minVal = c
	} else if c > maxVal {
		maxVal = c
	}
	return maxVal - minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
