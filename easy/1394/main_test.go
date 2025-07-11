package main

import "testing"

func Test1(t *testing.T) {
	got := findLucky([]int{2, 2, 3, 4})
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
