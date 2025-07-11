package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := permute([]int{1, 2, 3})
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
