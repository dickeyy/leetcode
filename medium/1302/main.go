package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 9.23mb (beats 36.61% of Go solutions)
 */

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := []*TreeNode{root}

	// process level by level
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		// process current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // dequeue

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		sum = levelSum
	}

	return sum
}
