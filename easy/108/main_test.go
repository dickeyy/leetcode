package main

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			name: "Test 1",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: &TreeNode{Val: 5}}},
		},
		{
			name: "Test 2",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: &TreeNode{Val: 5}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sortedArrayToBST(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
