package main

import "github.com/dickeyy/leetcode/utils"

func main() {
	utils.AssertEqual([][]int{{1, 3}, {3, 2}, {2, 1}}, validArrangement([][]int{{1, 3}, {3, 2}, {2, 1}}), 1)
}

/*
 *   Stats:
 *   Runtime: 361ms (beats 5.24% of Go solutions) :)
 *   Memory: 77.93mb (beats 5.07% of Go solutions) :(
 */

// Graph represents adjacency list and tracks in/out degrees
type Graph struct {
	adj      map[int][]int       // adjacency list
	inDeg    map[int]int         // in-degree count
	outDeg   map[int]int         // out-degree count
	edgeUsed map[int]map[int]int // track used edges
}

// validArrangement finds a valid arrangement of pairs where consecutive pairs share an element
func validArrangement(pairs [][]int) [][]int {
	// Initialize graph
	g := &Graph{
		adj:      make(map[int][]int),
		inDeg:    make(map[int]int),
		outDeg:   make(map[int]int),
		edgeUsed: make(map[int]map[int]int),
	}

	// Build graph
	for _, pair := range pairs {
		from, to := pair[0], pair[1]

		// Add edge to adjacency list
		g.adj[from] = append(g.adj[from], to)

		// Update degrees
		g.outDeg[from]++
		g.inDeg[to]++

		// Initialize edge usage tracking
		if g.edgeUsed[from] == nil {
			g.edgeUsed[from] = make(map[int]int)
		}
		g.edgeUsed[from][to]++
	}

	// Find start node (can be any node for Eulerian cycle)
	start := pairs[0][0]
	for node := range g.adj {
		// If out degree is one more than in degree, this must be start
		if g.outDeg[node]-g.inDeg[node] == 1 {
			start = node
			break
		}
	}

	// Run Hierholzer's algorithm to find Eulerian path
	path := []int{}
	hierholzer(g, start, &path)

	// Convert path back to pairs format
	result := make([][]int, len(path)-1)
	for i := len(path) - 1; i > 0; i-- {
		result[len(path)-1-i] = []int{path[i], path[i-1]}
	}

	return result
}

// hierholzer implements Hierholzer's algorithm to find Eulerian path
func hierholzer(g *Graph, node int, path *[]int) {
	// Process all outgoing edges
	for len(g.adj[node]) > 0 {
		// Get next neighbor
		next := g.adj[node][len(g.adj[node])-1]

		// Remove used edge
		g.adj[node] = g.adj[node][:len(g.adj[node])-1]
		g.edgeUsed[node][next]--

		// Recursively process neighbor
		hierholzer(g, next, path)
	}

	// Add current node to path
	*path = append(*path, node)
}
