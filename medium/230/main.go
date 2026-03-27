package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	a := []int{}
	curr := root
	for {
		a = append(a, curr.Val)
		if curr.Left != nil {
			curr = curr.Left
		} else {
			break
		}
	}
	slices.Sort(a)
	return a[k-1]
}
