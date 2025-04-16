package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := validArrangement([][]int{{1, 3}, {3, 2}, {2, 1}})
	want := [][]int{{1, 3}, {3, 2}, {2, 1}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
