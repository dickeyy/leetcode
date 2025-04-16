package main

import "testing"

func Test1(t *testing.T) {
	got := checkIfExist([]int{10, 2, 5, 3})
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
