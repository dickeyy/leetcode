package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := buildArray([]int{0, 2, 1, 5, 3, 4})
	want := []int{0, 1, 2, 4, 5, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
