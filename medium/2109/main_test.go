package main

import "testing"

func Test1(t *testing.T) {
	got := addSpaces("HiImHere", []int{2, 4, 8})
	want := "Hi Im Here"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
