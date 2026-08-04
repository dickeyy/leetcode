package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		got  []int
		want []int
	}{
		{
			desc: "test 1",
			got:  findMissingElements([]int{1, 4, 2, 5}),
			want: []int{3},
		},
		{
			desc: "test 2",
			got:  findMissingElements([]int{7, 8, 6, 9}),
			want: []int{},
		},
		{
			desc: "test 3",
			got:  findMissingElements([]int{5, 1}),
			want: []int{2, 3, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if !reflect.DeepEqual(tC.got, tC.want) {
				t.Errorf("got %v, want %v", tC.got, tC.want)
			}
		})
	}
}
