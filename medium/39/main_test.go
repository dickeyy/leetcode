package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	want := [][]int{{2, 2, 3}, {7}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
