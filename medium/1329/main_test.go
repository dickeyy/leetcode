package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}})
	want := [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
