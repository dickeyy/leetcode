package main

import "testing"

func Test1(t *testing.T) {
	got := sum(1, 1)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
